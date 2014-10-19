package main

import (
	"math"

	"github.com/mitsuse/progress-go"
)

const TASK_SIZE = 500
const BAR_WIDTH = 60

func main() {
	// Create a new progress bar.
	p := progress.NewSimple(TASK_SIZE, BAR_WIDTH)
	p.Show()

	for task := 0; task < p.Task(); task++ {
		doSomething()

		// Make a progress
		p.Add(1)
	}

	// Force to refresh the progress bar and quit.
	p.Close()
}

func doSomething() {
	iteration := int(math.Pow(10, 8))

	for i := 0; i < iteration; i++ {
	}
}
