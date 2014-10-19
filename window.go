package progress

import (
	"syscall"
	"unsafe"
)

type Window struct {
	rows uint16
	cols uint16
}

func (w *Window) Rows() int {
	return int(w.rows)
}

func (w *Window) Cols() int {
	return int(w.cols)
}

func GetWindow() *Window {
	valueSeq := [4]uint16{}

	result, _, errNumber := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&valueSeq)),
	)

	if int(result) == -1 {
		panic(errNumber)
	}

	w := &Window{
		rows: valueSeq[0],
		cols: valueSeq[1],
	}

	return w
}
