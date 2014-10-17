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
}

type progressBar struct {
	progress   int
	task       int
	tickCh     chan time.Time
	finalizeCh chan struct{}
}

func New(task int) ProgressBar {
	p := &progressBar{
		progress:   0,
		task:       task,
		tickCh:     make(chan time.Time),
		finalizeCh: make(chan struct{}),
	}

	return p
}

func (p *progressBar) Show() {
	p.refresh()

	ticker := time.NewTicker(time.Millisecond * 100)
	go func() {
		for t := range ticker.C {
			p.tickCh <- t
		}
	}()

	go func() {
		for _ = range p.tickCh {
			p.refresh()
		}

		ticker.Stop()
		p.finalize()
	}()
}

func (p *progressBar) Close() {
	close(p.tickCh)

	<-p.finalizeCh
}

func (p *progressBar) Add(progress int) {
	if update := p.progress + progress; update > p.task {
		p.progress = p.task
	} else {
		p.progress = update
	}
}

func (p *progressBar) refresh() {
	task := float64(p.task)
	progress := float64(p.progress)
	ratio := progress / task

	// TODO: Obtain the width of progress bar as a argument of New().
	progressStr := strings.Repeat("#", int(60*ratio))

	window := GetWindow()
	fmt.Print(strings.Repeat("\b", window.Cols()))

	fmt.Printf("\r%.1f%% %s", ratio*100, progressStr)
}

func (p *progressBar) finalize() {
	p.refresh()
	fmt.Print("\n")

	p.finalizeCh <- struct{}{}
}
