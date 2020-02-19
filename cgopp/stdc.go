package cgopp

/*
#include <string.h>
#include <stdlib.h>
#include <malloc.h>

*/
import "C"

import (
	"gopp"
	"reflect"
	"strings"
	"unsafe"
)

// std c library functions
// 这么封装一次，引用的包不需要再显式的引入"C"包。让CGO代码转换不传播到引用的代码中
func Cmemcpy()                 {}
func Cfree(ptr unsafe.Pointer) { C.free(ptr) }
func Cfree2(ptr *C.char)       { Cfree(unsafe.Pointer(ptr)) }
func Cfree3(ptrx interface{}) {
	// support types: uintptr, *foo._Ctype_char, unsafe.Pointer
	switch ptr := ptrx.(type) {
	case unsafe.Pointer:
		Cfree(ptr)
	case uintptr:
		Cfree(unsafe.Pointer(ptr))
	default:
		refty := reflect.TypeOf(ptrx)
		refval := reflect.ValueOf(ptrx)
		// *somepkg._Ctype_char
		if strings.Count(refty.String(), "*") > 0 &&
			strings.HasSuffix(refty.String(), "._Ctype_char") {
			ptr2 := unsafe.Pointer(refval.Pointer())
			Cfree(ptr2)
		} else {
			panic("unimpl " + refty.String())
		}
	}
}
func Calloc()   {}
func CMemset()  {}
func CMemZero() {}
func CStrcpy()  {}
func CStrdup()  {}

const CBoolTySz = gopp.Int32TySz
const CppBoolTySz = gopp.Int8TySz

// let freed memory really given back to OS
func MallocTrim() int { return int(C.malloc_trim(0)) }
