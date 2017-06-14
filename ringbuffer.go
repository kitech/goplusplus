package gopp

import (
	"time"

	"github.com/Workiva/go-datastructures/queue"
)

/*
用channel模拟的实现方式。
*/
type RingBuffer struct {
	n int
	c chan interface{}
}

func NewRingBuffer(n int) *RingBuffer {
	this := &RingBuffer{n: n}
	this.c = make(chan interface{}, n)
	return this
}

func (this *RingBuffer) Put(item interface{}) error {
	this.c <- item
	return nil
}

func (this *RingBuffer) Offer(item interface{}) (bool, error) {
	return false, nil
}

func (this *RingBuffer) Get() (item interface{}, err error) {
	item = <-this.c
	return
}

func (this *RingBuffer) Len() int {
	return len(this.c)
}

func (this *RingBuffer) Cap() int {
	return this.n
}

func (this *RingBuffer) Dispose() {
	return
}

func (this *RingBuffer) IsDisposed() bool {
	return false
}

//////
/*
封装的无锁rb，用一个比较丑的方式修正原来实现Get的100% CPU问题
*/
type RingBufferx struct {
	lfrb *queue.RingBuffer
}

func NewRingBufferx(n int) *RingBufferx {
	this := &RingBufferx{}
	this.lfrb = queue.NewRingBuffer(uint64(n))
	return this
}

func (this *RingBufferx) Put(item interface{}) error {
	return this.lfrb.Put(item)
}

func (this *RingBufferx) Offer(item interface{}) (bool, error) {
	return this.lfrb.Offer(item)
}

func (this *RingBufferx) Get() (item interface{}, err error) {
	for this.lfrb.Len() == 0 {
		time.Sleep(50 * time.Second)
	}
	return this.lfrb.Get()
}

func (this *RingBufferx) Len() uint64 {
	return this.lfrb.Len()
}

func (this *RingBufferx) Cap() uint64 {
	return this.lfrb.Cap()
}

func (this *RingBufferx) Dispose() { this.lfrb.Dispose() }

func (this *RingBufferx) IsDisposed() bool {
	return this.lfrb.IsDisposed()
}
