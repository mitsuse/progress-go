package progress

type ProgressBar interface {
	Progress() (progress uint)
	Task() (task uint)
	Show()
	Add(progress uint)
}

type progressBar struct {
	progress uint
	task     uint
}

func New(task uint) ProgressBar {
	p := &progressBar{
		progress: 0,
		task:     task,
	}

	return p
}

func (p *progressBar) Progress() (progress uint) {
	progress = p.progress
	return
}

func (p *progressBar) Task() (task uint) {
	task = p.task
	return
}

func (p *progressBar) Show() {
	// TODO: Implement this method.
}

func (p *progressBar) Add(progress uint) {
	// TODO: Implement this method.
}

func (p *progressBar) refresh() {
	// TODO: Implement this method.
}
