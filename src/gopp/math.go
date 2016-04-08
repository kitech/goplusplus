package gopp

import (
	"math"
)

func AbsI64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func AbsI32(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxU32(nums []uint32) uint32 {
	var ret uint32
	for _, n := range nums {
		if n > ret {
			ret = n
		}
	}
	return ret
}

func MinU32(nums []uint32) uint32 {
	var ret uint32 = math.MaxUint32
	for _, n := range nums {
		if n < ret {
			ret = n
		}
	}
	return ret
}
