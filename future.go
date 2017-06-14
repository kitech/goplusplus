package gopp

import (
	"context"
	"log"
)

type FutureFunc func(context.Context) (interface{}, error)
type Future struct {
	result    interface{}
	err       error
	completed bool
	ffunc     FutureFunc
	retC      chan struct{}
}

func NewFuture(ffunc FutureFunc) *Future {
	this := &Future{}
	this.ffunc = ffunc
	this.retC = make(chan struct{}, 1)

	// TODO cancel/timeout
	go func() {
		this.result, this.err = this.ffunc(nil)
		this.completed = true
		this.retC <- struct{}{}
	}()
	return this
}

func (this *Future) Get() (result interface{}, err error) {
	<-this.retC
	result = this.result
	err = this.err
	return
}
func (this *Future) GetOrTimeout() {}
func (this *Future) Cancel()       {}

func (this *Future) IsComplete() bool { return this.completed }

type Promise struct {
}

func (this *Promise) Then()           {}
func (this *Promise) WhenAll()        {}
func (this *Promise) WhenAny()        {}
func (this *Promise) WhenAnyMatched() {}
func (this *Promise) Pipe()           {}
func (this *Promise) Cancel()         {}
func (this *Promise) IsCanceled()     {}

type Lazyer interface {
	Eval() interface{}
}

/*
func Expand3TType(v []TType) {

}

func Expand3Int(v []int) (int, int, int) {
	return v[0], v[1], v[2]
}
*/

func init_future() {
	v := make([]int, 0)
	// v0,v1,v2 := v... // failed
	v0, v1, v2 := Expand3Int(v)
	if false {
		log.Println(v0, v1, v2)
	}
}

// not supported
//func get_varidic_results() (err error, rets...interface{}) {	return}
