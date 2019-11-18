package cgopp

/*
#include <string.h>
*/
import "C"
import "unsafe"

// 把浮点数存储在uint64中
func Float64AsInt(n float64) (rv uint64) {
	C.memcpy((unsafe.Pointer(&rv)), (unsafe.Pointer(&n)), 8)
	return
}
func Float32AsInt(n float32) (rv uint64) {
	C.memcpy((unsafe.Pointer(&rv)), (unsafe.Pointer(&n)), 4)
	return
}
func IntAsFloat64(v uint64) (n float64) {
	C.memcpy((unsafe.Pointer(&n)), (unsafe.Pointer(&v)), 8)
	return
}
func IntAsFloat32(v uint64) (n float32) {
	C.memcpy((unsafe.Pointer(&n)), (unsafe.Pointer(&v)), 4)
	return
}
func U64ToPtr(v uint64) unsafe.Pointer    { return unsafe.Pointer(uintptr(v)) }
func U64OfPtr(vptr unsafe.Pointer) uint64 { return uint64(uintptr(vptr)) }

func C2goBool(ok C.int) bool {
	if ok == 1 {
		return true
	}
	return false
}

func Go2cBool(ok bool) C.int {
	if ok {
		return 1
	}
	return 0
}

//
type go2cfnty *[0]byte

// 参数怎么传递
func Go2cfnp(fn unsafe.Pointer) *[0]byte {
	return go2cfnty(fn)
}
func Go2cfn(fn interface{}) *[0]byte {
	// assert(reflect.TypeOf(fn).Kind == reflect.Ptrx)
	return Go2cfnp(fn.(unsafe.Pointer))
}
