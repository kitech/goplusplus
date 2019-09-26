package cgopp

/*
#include <string.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

// x64
// note: C.int != go int
type Cint C.int
type Cgoint int32
type Clong C.long
type Cgolong int64

// x32
// note: C.int == go int
/*
type Cint C.int
type Cgoint int32
type Clong C.long
type Cgolong = int64
*/

// => char**
type CStrArr struct {
	carr  unsafe.Pointer
	calen int
	garr  []*string
}

func CStrArrFromu8(arr **uint8, n int) *CStrArr {
	return CStrArrFromp(unsafe.Pointer(arr), n)
}
func CStrArrFromc8(arr **int8, n int) *CStrArr {
	return CStrArrFromp(unsafe.Pointer(arr), n)
}

// must a (u)char**
func CStrArrFromp(arr unsafe.Pointer, n int) *CStrArr {
	this := &CStrArr{}
	this.carr = arr
	return this
}

func (this *CStrArr) ToGo() (rets []string) {
	for i := 0; i < this.calen; i++ {
		ep := unsafe.Pointer(uintptr(this.carr) + unsafe.Sizeof((uintptr(0)))*uintptr(i))
		e := C.GoString((*C.char)(ep))
		rets = append(rets, e)
	}
	return
}

func CStrArrFromStrs(arr []string) *CStrArr {
	this := &CStrArr{}
	for _, e := range arr {
		t := e + "\x00"
		this.garr = append(this.garr, &t)
	}
	this.garr = append(this.garr, nil)
	return this
}

func (this *CStrArr) ToC() unsafe.Pointer {
	return (unsafe.Pointer)(&this.garr[0])
}

func (this *CStrArr) Append(s string) {
	if this.garr == nil {
		// think as from c
		strs := this.ToGo()
		tarr := CStrArrFromStrs(strs)
		this.garr = tarr.garr
	}
	e := s + "\x00"
	this.garr = append(this.garr, &e)
}

func GoStrArr2c(arr []string) uintptr {
	if len(arr) == 0 {
		return 0
	}

	pv := make([]unsafe.Pointer, len(arr)+1)
	for i, v := range arr {
		pv[i] = unsafe.Pointer(C.CString(v))
	}
	sz := int(unsafe.Sizeof(uintptr(0))) * (len(arr) + 1)
	rv := C.calloc(1, C.ulong(sz))
	C.memcpy(rv, unsafe.Pointer(&pv[0]), C.ulong(sz))
	return uintptr(rv)
}
