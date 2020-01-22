package xlog

/*
#include <stdio.h>
#include <stdarg.h>
*/
import "C"
import "gopp/xstrings"

const skip_depth = 4
const pkgsep = "__"
const mthsep = "_"

var gopaths []string

func init() {
	eptr := C.getenv("GOPATH".ptr)
	estr := C.GoString(eptr)
	gopaths = xstrings.Split(estr, ":")

	assert(len(pkgsep) > 0)
}

func dummy(args ...interface{}) {

}

func Println(args ...interface{}) {
	for i := 0; i < len(args); i++ {

	}
	// dummy(args...) // not work
}

func Fatalln(args ...interface{}) {

}

func Printf(format string, args ...interface{}) {

}

func trim_gopath(s string) string {
	const clen = 5 // "/src/" length
	for i := 0; i < len(gopaths); i++ {
		sub := gopaths[i]
		if xstrings.Prefixed(s, sub) {
			return s[len(sub)+clen:]
		}
	}
	return s
}

func demangle_funcname(s string) string {
	s2 := xstrings.Replace(s, pkgsep, ".", 1)
	return s2
}

// TODO
func printfmt(format string, a0 interface{}) {
	// tyid := a0.tyid
}
func printint(a0 int) {
	callers := Callers()
	assert(len(callers) > skip_depth)
	caller := callers[skip_depth]
	// println(caller.File, ":", caller.Lineno, caller.Funcname, a0)
	trfile := trim_gopath(caller.File)
	funcname := demangle_funcname(caller.Funcname)
	C.printf("%s:%d %s %d\n".ptr, trfile.ptr, caller.Lineno, funcname.ptr, a0)
}
func printstr(a0 string) {
	callers := Callers()
	assert(len(callers) > skip_depth)
	caller := callers[skip_depth]
	// println(caller.File, ":", caller.Lineno, caller.Funcname, a0)
	trfile := trim_gopath(caller.File)
	funcname := demangle_funcname(caller.Funcname)
	C.printf("%s:%d %s %.*s\n".ptr,
		trfile.ptr, caller.Lineno, funcname.ptr, a0.len, a0.ptr)
}
func printptr(a0 voidptr) {
	callers := Callers()
	assert(len(callers) > skip_depth)
	caller := callers[skip_depth]
	// println(caller.File, ":", caller.Lineno, caller.Funcname, a0)
	trfile := trim_gopath(caller.File)
	funcname := demangle_funcname(caller.Funcname)
	C.printf("%s:%d %s %p\n".ptr, trfile.ptr, caller.Lineno, funcname.ptr, a0)
}
func printflt(a0 f64) {
	callers := Callers()
	assert(len(callers) > skip_depth)
	caller := callers[skip_depth]
	// println(caller.File, ":", caller.Lineno, caller.Funcname, a0)
	trfile := trim_gopath(caller.File)
	funcname := demangle_funcname(caller.Funcname)
	C.printf("%s:%d %s %g\n".ptr, trfile.ptr, caller.Lineno, funcname.ptr, a0)
}

func Keep() {}
