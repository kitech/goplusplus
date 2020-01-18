package xos

/*
#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
*/
import "C"

// import "unsafe"

func Touch(path string) bool {
	// fp := C.open(path.ptr, C.O_RDWR|C.O_CREAT)
	C.close(-1)
	return true
}

func Environ() {
	a := C.__environ
	a0 := a[0]
	a00 := a[0][0]
	b := C.O_RDWR
	c := C.int(1)
	println(a0)
	println(a00)
}

func Keep() {}
