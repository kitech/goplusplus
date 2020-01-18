package xlog

/*
#cgo LDFLAGS: -ldwarf

#include <execinfo.h>
#include <libdwarf/dwarf.h>
#include <libdwarf/libdwarf.h>
#include <libelf.h>

*/
import "C"

type Frame struct {
	btdepth int

	funcname string
	mglname  string
	// funcaddr unsafe.Pointer
	file string
	line string

	sframe string
}

func Backtrace() {
	// var buf = make([]byte, 100)
	buf1 := []byte{}
	buf := C.cxmalloc(200)
	nr := C.backtrace(buf, 200/8)
	// println("nr=", nr)
	symarr := C.backtrace_symbols(buf, nr)
	defer C.free(symarr)
	for i := 0; i < nr; i++ {
		symit := symarr[i]
		// symstr := string(symit)
		// symcstr := (*C.char)(symit)
		symstr := C.GoString(symit)
		println(symit)
		println(symstr)
	}
}
