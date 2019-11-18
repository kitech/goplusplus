package gopp

import (
	"log"
	"math"
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

/////
// see funk.Chunk
func ChunkN(arr interface{}, c int) interface{} {
	arrv := reflect.ValueOf(arr)
	n := int(math.Ceil(float64(arrv.Len()) / float64(c)))
	_ = n
	return nil
}

// safe version
func SliceGetStr(arr []string, idx int, dft string) string {
	if len(arr) > idx {
		return arr[idx]
	}
	return dft
}

func IntsSame(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	same := true
	for i := 0; i < len(arr)-1; i++ {
		same = same && arr[i] == arr[i+1]
	}
	return same
}
func StrsSame(arr []string) bool {
	if len(arr) <= 1 {
		return false
	}
	same := true
	for i := 0; i < len(arr)-1; i++ {
		same = same && arr[i] == arr[i+1]
	}
	return same
}

// 等差，且差不等于0
func IntsDiffeq(arr []int) bool {
	if len(arr) <= 1 {
		return false
	}
	same := true
	diff := arr[1] - arr[0]
	if diff == 0 {
		return false
	}
	for i := 0; i < len(arr)-1; i++ {
		diff2 := arr[i+1] - arr[i]
		same = same && diff2 == diff
	}
	return same
}
