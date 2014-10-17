package main

import (
	"math"

	"github.com/mitsuse/progress-go"
)

const TASK_SIZE = 500
const WORKER_NUM = 2

func main() {
	// Create a new progress bar.
	p := progress.New(TASK_SIZE)
	p.Show()

	quitCh := make(chan struct{}, WORKER_NUM)

	for i := 0; i < cap(quitCh); i++ {
		go func() {
			taskSize := TASK_SIZE / WORKER_NUM

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
