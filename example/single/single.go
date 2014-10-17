package main

import (
	"time"

	"github.com/mitsuse/progress-go"
)

const TASK_SIZE = 500

func main() {
	// Create a new progress bar.
	p := progress.New(TASK_SIZE)
	p.Show()

	for task := 0; task < TASK_SIZE; task++ {
		// Do something.
		time.Sleep(time.Millisecond * 10)

		// Make a progress
		p.Add(1)
	}

	// Force to refresh the progress bar and quit.
	p.Close()
}
