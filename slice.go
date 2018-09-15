package gopp

import (
	"log"
	"reflect"
	"unsafe"
)

// 一个vector/slice中是否包含一个值
func Vcontains(v interface{}, e interface{}) bool {
	ret := false

	fr := func(i interface{}) bool {
		if i == e {
			ret = true
			return true
		}
		return false
	}
	fo := MustFunc(fr)

	res := MapAny(fo, v)
	log.Println(res)

	return ret
}

// 一个 map中是否包含一个值
func Mcontains(m interface{}, e interface{}) bool {
	fr := func(i interface{}, iv interface{}) interface{} {
		return i == e || iv.(bool)
	}
	fo := MustFunc(fr)
	ret := ReduceAny(fo, m, false)
	return ret.(bool)
}

type Vec struct {
	v interface{}
}

func NewVec(v interface{}) *Vec { return &Vec{v} }

type Set struct{}

type Tuple struct {
}

type Triple struct {
}

// also need a copy, if not modify original one
func SliceTyConv(src []int16) []byte {
	src_ := make([]int16, len(src))
	copy(src_, src)
	srchdr := (*reflect.SliceHeader)(unsafe.Pointer(&src_[0]))
	srchdr.Len = len(src_) * 2 // int(unsafe.Sizeof(int16(0))/unsafe.Sizeof(byte(0)))
	srchdr.Cap = cap(src_) * 2 //int(unsafe.Sizeof(int16(0))/unsafe.Sizeof(byte(0)))
	return *((*[]byte)(unsafe.Pointer(srchdr)))
}
