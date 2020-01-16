package xos

/*
#include <errno.h>
#include <string.h>
*/
import "C"

func Errno() int {
	// because cgo cannot refer to errno directly, so use another way
	return *C.__errno_location()
}

func Errmsg() string {
	eno := Errno()
	emsg := C.strerror(eno)
	return string(emsg)
}
