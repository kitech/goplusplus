package cgopp

/*
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"unsafe"

	"gopp"
)

// 需要关闭的对象的自动处理
// *os.File, *http.Response
func Deferx(objx interface{}) {
	if objx == nil {
		return
	}

	switch obj := objx.(type) {
	case *C.char:
		runtime.SetFinalizer(objx, func(objx interface{}) { C.free(unsafe.Pointer(obj)) })
	default:
		gopp.Deferx(objx)
	}
}
