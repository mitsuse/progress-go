# Progress

[![Build Status](https://travis-ci.org/mitsuse/progress-go.svg?branch=master)](https://travis-ci.org/mitsuse/progress-go)
[![GoDoc](http://godoc.org/github.com/mitsuse/progress-go?status.svg)](http://godoc.org/github.com/mitsuse/progress-go)

A library to show progressbars in CLI implemented in Golang. 

## Example

```go
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
```


## TODO

- Implement custumizable progressbar.
- Use channel to update the progress and the bar.
- Write the docs.

## License

The license is not decided yet.
