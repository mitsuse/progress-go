# Progress

[![Build Status](https://travis-ci.org/mitsuse/progress-go.svg?branch=master)](https://travis-ci.org/mitsuse/progress-go)
[![GoDoc](http://godoc.org/github.com/mitsuse/progress-go?status.svg)](http://godoc.org/github.com/mitsuse/progress-go)

A library to show customizable progress bar on CLI implemented in Golang.

## Features

Porgress-go supports the following features:

- Custamizable widgets
- Status update with multi-threads

## Example

The following snippet is an example of simple progress bar.

```go
// Create a new progress bar. 
// "TASK_SIZE" is the total number of tasks to be processed.
// "BAR_WIDTH" is the width of progress bar.
p := progress.NewSimple(TASK_SIZE, BAR_WIDTH)
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

- Write the docs.
- Support for various platforms.

## License

The license is not decided yet.
