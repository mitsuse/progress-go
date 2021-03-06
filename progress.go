package progress

import (
	"fmt"
	"strings"
	"time"
)

type ProgressBar interface {
	Show()
	Close()
	Add(progress int)
	Task() int
}

type progressBar struct {
	progress  int
	task      int
	startTime int64
	ticker    *time.Ticker

	stopCh     chan struct{}
	finalizeCh chan struct{}

	addSynCh chan struct{}

	widgetSeq []Widget
}

func New(task int, widgetSeq []Widget) ProgressBar {
	p := &progressBar{
		progress: 0,
		task:     task,

		stopCh:     make(chan struct{}),
		finalizeCh: make(chan struct{}),

		addSynCh: make(chan struct{}, 1),

		widgetSeq: widgetSeq,
	}

	return p
}

func NewSimple(task, width int) ProgressBar {
	widgetSeq := []Widget{
		NewPercentage("%.1f%%"),
		NewBar(width, "#", ""),
	}

	return New(task, widgetSeq)
}

func (p *progressBar) Show() {
	p.startTime = time.Now().Unix()

	p.refresh()

	ticker := time.NewTicker(time.Millisecond * 100)
	go func() {

		for running := true; running; {
			select {
			case <-ticker.C:
				p.refresh()
			case <-p.stopCh:
				ticker.Stop()
				running = false
			}
		}

		p.finalize()
	}()
}

func (p *progressBar) Close() {
	p.stopCh <- struct{}{}
	<-p.finalizeCh
}

func (p *progressBar) Add(progress int) {
	p.addSynCh <- struct{}{}

	if update := p.progress + progress; update > p.task {
		p.progress = p.task
	} else {
		p.progress = update
	}

	<-p.addSynCh
}

func (p *progressBar) Task() int {
	return p.task
}

func (p *progressBar) refresh() {
	size, err := GetSize()
	if err != nil {
		return
	}

	fmt.Print(strings.Repeat("\b", size.Cols()))
	fmt.Print("\r")

	if len(p.widgetSeq) > 0 {
		fmt.Print(p.widgetSeq[0].Print(p.progress, p.task, p.startTime))
	}

	for i := 1; i < len(p.widgetSeq); i++ {
		widget := p.widgetSeq[i]
		fmt.Printf(" %s", widget.Print(p.progress, p.task, p.startTime))
	}
}

func (p *progressBar) finalize() {
	p.refresh()
	fmt.Print("\n")

	p.finalizeCh <- struct{}{}
}
