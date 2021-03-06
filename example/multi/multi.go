package main

import (
	"math"

	"github.com/mitsuse/progress-go"
)

const TASK_SIZE = 500
const BAR_WIDTH = 60
const WORKER_NUM = 2

func main() {
	// Create a new progress bar.
	p := progress.NewSimple(TASK_SIZE, BAR_WIDTH)
	p.Show()

	quitCh := make(chan struct{}, WORKER_NUM)

	for i := 0; i < cap(quitCh); i++ {
		go func() {
			taskSize := p.Task() / cap(quitCh)

			for x := 0; x < taskSize; x++ {
				doSomething()

				// Make a progress
				p.Add(1)
			}

			quitCh <- struct{}{}
		}()
	}

	// Wait for workers to finish.
	for i := 0; i < cap(quitCh); i++ {
		_ = <-quitCh
	}

	// Force to update and quit the progress bar.
	p.Close()
}

func doSomething() {
	iteration := int(math.Pow(10, 8))

	for i := 0; i < iteration; i++ {
	}
}
