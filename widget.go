package progress

import (
	"fmt"
	"math"
	"strings"
)

type Widget interface {
	Print(progress, task int) (repr string)
}

type Perpcentage struct {
	format string
}

func NewPercentage(format string) *Perpcentage {
	p := &Perpcentage{
		format: format,
	}

	return p
}

func (p *Perpcentage) Print(progress, task int) (repr string) {
	return fmt.Sprintf(p.format, float64(progress)/float64(task)*100)
}

type Bar struct {
	size     int
	progress string
	rest     string
}

func NewBar(size int, progress, rest string) *Bar {
	b := &Bar{
		size:     size,
		progress: progress,
		rest:     rest,
	}

	return b
}

func (b *Bar) Print(progress, task int) (repr string) {
	template := "%s%s"

	progressSize := int(math.Floor(float64(progress) / float64(task) * float64(b.size)))
	progressBar := strings.Repeat(b.progress, progressSize)

	restSize := b.size - progressSize
	restBar := strings.Repeat(b.rest, restSize)

	return fmt.Sprintf(template, progressBar, restBar)
}
