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
