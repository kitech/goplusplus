package xos

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
*/
import "C"

// import "unsafe"

const (
	DftMode  = 0644
	DftMask  = 0022
	PATH_MAX = 256 // C.PATH_MAX
)

func Touch(path string) bool {
	fp := C.open(path.ptr, C.O_RDWR|C.O_CREAT, 0644)
	if fp < 0 {
		return false
	}
	C.close(fp)
	return true
}

func Environ() []string {
	arr := []string{}
	envp := C.__environ
	envp0 := envp[0]
	envp00 := envp[0][0]
	b := C.O_RDWR
	c := C.int(1)
	println(envp0)
	println(envp00)
	return arr
}

func Keep() {}
