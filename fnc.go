package gopp

/*
#include <stdint.h>
*/
import "C"
import (
	"log"
	"reflect"
	"sync"
	"unsafe"
)

//export InvokeGoFunc
func InvokeGoFunc(fp unsafe.Pointer, argc C.int, argv []C.uint64_t) {

}

type Func4cb struct {
	Fp      unsafe.Pointer
	Convers []func(uint64) interface{}
}
type F4cConvType func(uint64) interface{}

func (this *Func4cb) CPtr() unsafe.Pointer { return unsafe.Pointer(this) }

func NewFunc4cbFrom(fp unsafe.Pointer) *Func4cb { return (*Func4cb)(fp) }
func NewFunc4cb(fp unsafe.Pointer, cvs []func(uint64) interface{}) *Func4cb {
	this := &Func4cb{}
	this.Fp = fp
	this.Convers = cvs
	return this
}

// hold reference
type f4ccfunctors struct {
	mu  sync.Mutex
	fts map[unsafe.Pointer]*Func4cb
}

var f4ccrefs = &f4ccfunctors{fts: make(map[unsafe.Pointer]*Func4cb)}

func F4ccRef(ft *Func4cb) {
	f4ccrefs.mu.Lock()
	defer f4ccrefs.mu.Unlock()

	if _, ok := f4ccrefs.fts[ft.Fp]; !ok {
		f4ccrefs.fts[ft.Fp] = ft
	}
}

func F4ccUnref(ft *Func4cb) {
	f4ccrefs.mu.Lock()
	defer f4ccrefs.mu.Unlock()

	if _, ok := f4ccrefs.fts[ft.Fp]; ok {
		delete(f4ccrefs.fts, ft.Fp)
	}
}

//export InvokeGoFuncWithArgConver
func InvokeGoFuncWithArgConver(ft unsafe.Pointer, argc C.int, argv []C.uint64_t) {
	fp := (*Func4cb)(ft)
	log.Println(fp)

	argv_ifs := []interface{}{}
	for idx, cv := range fp.Convers {
		argv_ifs = append(argv_ifs, cv(uint64(argv[idx])))
	}

	argv_vals := []reflect.Value{}
	for _, argx := range argv_ifs {
		argv_vals = append(argv_vals, reflect.ValueOf(argx))
	}

	fpx := (interface{})(fp.Fp)
	fpval := reflect.ValueOf(fpx)
	retvals := fpval.Call(argv_vals)
	_ = retvals
}

func F4ccArgConvInt(arg uint64) interface{}   { return int(arg) }
func F4ccArgConvInt8(arg uint64) interface{}  { return int8(arg) }
func F4ccArgConvInt16(arg uint64) interface{} { return int16(arg) }
func F4ccArgConvInt32(arg uint64) interface{} { return int32(arg) }
func F4ccArgConvInt64(arg uint64) interface{} { return int64(arg) }

func F4ccArgConvUInt(arg uint64) interface{}   { return uint(arg) }
func F4ccArgConvUInt8(arg uint64) interface{}  { return uint8(arg) }
func F4ccArgConvUInt16(arg uint64) interface{} { return uint16(arg) }
func F4ccArgConvUInt32(arg uint64) interface{} { return uint32(arg) }
func F4ccArgConvUInt64(arg uint64) interface{} { return uint64(arg) }

func F4ccArgConvFloat(arg uint64) interface{}   { return float32(arg) }
func F4ccArgConvFloat64(arg uint64) interface{} { return float64(arg) }

func F4ccArgConvString(arg uint64) interface{} {
	argp := (*C.char)(unsafe.Pointer(uintptr(arg)))
	return C.GoString(argp)
}

func F4ccArgConvPointer(arg uint64) interface{} {
	return unsafe.Pointer(uintptr(arg))
}
