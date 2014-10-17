package main

import (
	"time"

	"github.com/mitsuse/progress-go"
)

func main() {
	// Create a new progress bar.
	// "500" is the size of task related to the progress bar.
	p := progress.New(500)
	p.Show()

	for task := 0; task < p.Task(); task++ {
		// Do something.
		time.Sleep(time.Millisecond * 10)

		// Make a progress
		p.Add(1)
	}

	// Force to update and quit the progress bar.
	p.Close()
}
