package gopp

import (
	"sync/atomic"
)

// TODO CPU Pause
// ABA问题预防 - DCAS
// note: not copy
type AtomicU32 uint64 // | 32b abaver | 32b realval |

func (this AtomicU32) CmpAndSwap(old uint32, new uint32) (swapped bool) {
	curv := atomic.LoadUint64((*uint64)(&this))
	abaver := uint32(curv >> 32)
	oldv := uint64(abaver)<<32 | uint64(old)
	newv := uint64(abaver+1)<<32 | uint64(new)

	return atomic.CompareAndSwapUint64((*uint64)(&this), oldv, newv)
}

type AtomicBool uint32

func NewAtomicBool() *AtomicBool {
	var this AtomicBool
	return &this
}

func (this *AtomicBool) CmpAndSwap(old bool, new bool) (swapped bool) {
	var oldiv, newiv uint32
	if old {
		oldiv = 1
	}
	if new {
		newiv = 1
	}

	curv := atomic.LoadUint32((*uint32)(this))
	abaver := curv >> 31
	oldv := abaver<<31 | oldiv
	newv := (abaver+1)<<31 | newiv

	swapped = atomic.CompareAndSwapUint32((*uint32)(this), oldv, newv)
	return
}

func (this *AtomicBool) IsTrue() bool {
	curv := atomic.LoadUint32((*uint32)(this))
	return curv&1 == 1
}

func (this *AtomicBool) Value() uint32 { return atomic.LoadUint32((*uint32)(this)) }

// https://www.jianshu.com/p/72d02353dc7e
//
