package progress

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Size struct {
	rows uint16
	cols uint16
}

func (s *Size) Rows() int {
	return int(s.rows)
}

func (s *Size) Cols() int {
	return int(s.cols)
}

func GetSize() (s *Size, err error) {
	valueSeq := [4]uint16{}

	result, _, errNumber := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&valueSeq)),
	)

	if int(result) == -1 {
		err = SizeError(errNumber)
	} else {
		s = &Size{
			rows: valueSeq[0],
			cols: valueSeq[1],
		}
	}

	return
}

type SizeError uintptr

func (e SizeError) Error() string {
	template := "Fail to get the size of terminal: %v"

	return fmt.Sprintf(template, e)
}
