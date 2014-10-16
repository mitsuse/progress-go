package progress

import (
	"fmt"
	"strings"
)

type ProgressBar interface {
	Progress() (progress uint)
	Task() (task uint)
	Show()
	Close()
	Add(progress uint)
	isShown() (shown bool)
}

type progressBar struct {
	progress   uint
	task       uint
	lastUpdate uint
	shown      bool
}

func New(task uint) ProgressBar {
	p := &progressBar{
		progress: 0,
		task:     task,
	}

	return p
}

func (p *progressBar) Progress() (progress uint) {
	return p.progress
}

func (p *progressBar) Task() (task uint) {
	task = p.task
	return
}

func (p *progressBar) Show() {
	p.shown = true

	p.refresh()
}

func (p *progressBar) Close() {
	if !p.isShown() {
		return
	}

	p.refresh()

	p.shown = false
	fmt.Print("\n")
}

func (p *progressBar) Add(progress uint) {
	if update := p.progress + progress; update > p.task {
		p.progress = p.task
	} else {
		p.progress = update
	}

	p.refresh()
}

func (p *progressBar) refresh() {
	if !p.isShown() {
		return
	}

	task := float64(p.Task())
	progress := float64(p.Progress())
	ratio := progress / task

	// TODO: Obtain the width of progress bar as a argument of New().
	progressStr := strings.Repeat("#", int(60*ratio))

	window := GetWindow()
	fmt.Print(strings.Repeat("\b", window.Cols()))

	fmt.Printf("\r%.1f%% %s", ratio*100, progressStr)
}

func (p *progressBar) isShown() (shown bool) {
	return p.shown
}
