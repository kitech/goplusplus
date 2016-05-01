package gopp

import (
	"reflect"
)

type FuncTools struct {
	first     chan interface{}
	last      chan interface{}
	tmpchains []chan interface{}
}

type Func struct {
	in  reflect.Type
	out reflect.Type
}

func Abc(T1 reflect.Type, T2 reflect.Type) {
	var n T1
	var a = 1
	var t3 reflect.Type = reflect.TypeOf(a)
	var b t33
}

func NewFuncTools(srcpeer chan interface{}) *FuncTools {
	first := srcpeer
	last := first
	tmpchains := make([]chan interface{}, 0)
	return &FuncTools{first, last, tmpchains}
}

func (this *FuncTools) Iter() chan interface{} {
	return this.last
}

func (this *FuncTools) FcMap(f func(interface{}) interface{}) *FuncTools {
	ch := make(chan interface{}, 0)
	curch := this.last
	this.last = ch
	this.tmpchains = append(this.tmpchains, ch)

	go func() {
		for e := range curch {
			ne := f(e)
			this.last <- ne
		}
		close(this.last)
	}()
	return this
}

func (this *FuncTools) FcFilter(f func(interface{}) bool) *FuncTools {
	ch := make(chan interface{}, 0)
	curch := this.last
	this.last = ch
	this.tmpchains = append(this.tmpchains, ch)

	go func() {
		for e := range curch {
			if f(e) {
				this.last <- e
			}
		}
		close(this.last)
	}()
	return this
}

func Fcmap(sets chan string, f func(string) bool) <-chan string {
	ch := make(chan string, 0)
	go func() {
		for e := range sets {
			if f(e) {
				ch <- e
			}
		}
		close(ch)
	}()

	return ch
}
