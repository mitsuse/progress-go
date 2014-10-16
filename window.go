package progress

import (
	"syscall"
	"unsafe"
)

type Window struct {
	rows   uint16
	cols   uint16
	height uint16
	width  uint16
}

func (w *Window) Rows() int {
	return int(w.rows)
}

func (w *Window) Cols() int {
	return int(w.cols)
}

func (w *Window) Height() int {
	return int(w.height)
}

func (w *Window) Width() int {
	return int(w.width)
}

func GetWindow() *Window {
	w := &Window{}

	result, _, errNumber := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(w)),
	)

	if int(result) == -1 {
		panic(errNumber)
	}

	return w
}
