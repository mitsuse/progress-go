package progress

import (
	"fmt"
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

func GetWindow() (w *Window, err error) {
	valueSeq := [4]uint16{}

	result, _, errNumber := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&valueSeq)),
	)

	if int(result) == -1 {
		err = WindowError(errNumber)
	} else {
		w = &Window{
			rows: valueSeq[0],
			cols: valueSeq[1],
		}
	}

	return
}

type WindowError uintptr

func (e WindowError) Error() string {
	template := "Fail to get the size of terminal: %v"

	return fmt.Sprint(template, e)
}
