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

for task := 0; task < TASK_SIZE; task++ {
	// Do something.
	time.Sleep(time.Millisecond * 10)

	// Make a progress
	p.Add(1)
}

// Force to refresh the progress bar and quit.
p.Close()
```

## TODO

- Implement custumizable progressbar.
- Write the docs.

## License

The license is not decided yet.
