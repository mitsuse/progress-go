# Progress

[![Build Status](https://travis-ci.org/mitsuse/progress-go.svg?branch=master)](https://travis-ci.org/mitsuse/progress-go)
[![GoDoc](http://godoc.org/github.com/mitsuse/progress-go?status.svg)](http://godoc.org/github.com/mitsuse/progress-go)

A library to show progressbars in CLI implemented in Golang. 

## Example

The following snippet is an example of simple progress bar.

```go
// Create a new progress bar. 
p := progress.New(TASK_SIZE) // TASK_SIZE = 500
p.Show()

for task := 0; task < p.Task(); task++ {
    doSomething()

	// Make a progress
	p.Add(1)
}

// Force to refresh the progress bar and quit.
p.Close()
```

Progress-go also works on multiple threads.
Please try running [multi.go](example/multi/multi.go).

## Installation

Just execute the following command:

```sh
go get github.com/mitsuse/progress-go
```

## TODO

- Implement custumizable progressbar.
- Write the docs.

## License

The license is not decided yet.
