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
	progress int
	task     int
	ticker   *time.Ticker

	stopCh     chan struct{}
	finalizeCh chan struct{}

	addSynCh chan struct{}

	widgetSeq []Widget
}

func New(task int) ProgressBar {
	// TODO: Obtain the width of progress bar as a argument of New().
	widgetSeq := []Widget{
		NewPercentage("%.1f%%"),
		NewBar(60, "#", ""),
	}

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

func (p *progressBar) Show() {
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
		fmt.Print(p.widgetSeq[0].Print(p.progress, p.task))
	}

	for i := 1; i < len(p.widgetSeq); i++ {
		widget := p.widgetSeq[i]
		fmt.Printf(" %s", widget.Print(p.progress, p.task))
	}
}

func (p *progressBar) finalize() {
	p.refresh()
	fmt.Print("\n")

	p.finalizeCh <- struct{}{}
}
