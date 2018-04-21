package gopp

type BitSet struct{ bits uint64 }

func NewBitSet() *BitSet               { return &BitSet{} }
func (this *BitSet) Add(i uint64)      { this.bits |= 1 << i }
func (this *BitSet) Del(i uint64)      { this.bits &= ^(1 << i) }
func (this *BitSet) Has(i uint64) bool { return this.bits&(1<<i) == 1 }
