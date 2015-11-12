package gopp

import (
	"fmt"
	"sync"
)

// 通用的可自动回收关闭的go channel封装
// 这个类型的封装不行啊，根本行不通，限制太多了。
var gwg sync.WaitGroup

type xchan struct {
	uchan  chan interface{}
	wg     sync.WaitGroup
	pdchan chan struct{}
	cdchan chan struct{}
	// 使用标识的话，就怕在不同goroutine使用时有问题。
}

//
func NewXChan(c int) *xchan {
	uchan := make(chan interface{}, c)
	pdchan := make(chan struct{}, 1)
	cdchan := make(chan struct{}, 1)

	var wg sync.WaitGroup
	wg.Add(1)

	gwg.Add(1)
	return &xchan{uchan, wg, pdchan, cdchan}
}

// Producer done
func (this *xchan) PDone() {
	if len(this.pdchan) > 0 {
		return
	}
	this.pdchan <- struct{}{}

	go func() {
		// 写入数据完成，读取数据完成
		this.wg.Wait()
		close(this.pdchan)
		close(this.cdchan)
		close(this.uchan)
		fmt.Println("done....")
		gwg.Done()
	}()
}

// Consumer done
func (this *xchan) CDone() {
	if len(this.cdchan) > 0 {
		return
	}
	this.cdchan <- struct{}{}

	this.wg.Done()
	// this.wg.Done()  // 怎么就不能检测一下呢
}

func (this *xchan) Write(v interface{}) {
	this.uchan <- v
}

func (this *xchan) Read() interface{} {
	v := <-this.uchan
	return v
}

func xchan_wait() {
	gwg.Wait()
}
