package progress

import (
	"fmt"
	"strings"
	"time"
)

type ProgressBar interface {
	Progress() (progress int)
	Task() (task int)
	Show()
	Close()
	Add(progress int)
	isShown() (shown bool)
}

type progressBar struct {
	progress   int
	task       int
	shown      bool
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

func (p *progressBar) Progress() (progress int) {
	return p.progress
}

func (p *progressBar) Task() (task int) {
	task = p.task
	return
}

func (p *progressBar) Show() {
	p.shown = true
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
	if !p.isShown() {
		return
	}
	p.shown = false

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
	task := float64(p.Task())
	progress := float64(p.Progress())
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

func (p *progressBar) isShown() (shown bool) {
	return p.shown
}
