package gopp

import "log"

type FuncTools struct {
	first     chan interface{}
	last      chan interface{}
	tmpchains []chan interface{}
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

// curry function，
// 爲了通用一點，可能涉及到interface{}
// 這只是個原地使用的例子
func applyFnDemo() {
	curry_src_fn := func(a int, b int) {}
	curry_dest_fn := func(b int) { curry_src_fn(123, b) }
	log.Println(&curry_dest_fn)
}

// see https://github.com/choleraehyq/gofunctools/functools
