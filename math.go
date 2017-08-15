package gopp

import (
	"log"
	"math"
)

func AbsNum(x interface{}) interface{} {
	f1 := func(x int64) int64 {
		if x < 0 {
			return -x
		}
		return x
	}
	f2 := func(x float64) float64 {
		if x < 0.0 {
			return -x
		}
		return x
	}

	fns := []interface{}{f1, f2}
	fnidx := SymbolResolveFns([]interface{}{x}, fns)
	if fnidx == -1 {
		log.Panicln("Unresolved")
	}

	out := overloadRcCall([]interface{}{x}, fns[fnidx])
	return out[0].Interface()
}

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
