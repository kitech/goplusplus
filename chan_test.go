package gopp

import (
	"log"
	"testing"
	"time"
)

/*
write channel close or timeout
*/

func TestChannel1(t *testing.T) {
	ch := make(chan bool)
	sendok := true
	sendval := true

	defer func() {
		if x := recover(); x != nil {
			sendok = false
			log.Printf("wow should be closed channel: %v", x)
		}
	}()

	timeoutSend := func() {
		select {
		case ch <- sendval:
		default:
			log.Println("send busch blocked:", len(ch))
			// TODO 这种情况是为什么呢，应该怎么办呢？
			// debug1.PrintStack()
			tmer := time.AfterFunc(5*time.Second, func() {
				panic("send busch timeout")
			})
			ch <- sendval
			stopok := tmer.Stop()
			if !stopok {
				panic("stop timer failed")
			}
		}
	}
	go timeoutSend()
	select {}
}

func TestC3(t *testing.T) {
	// 返回读取channel
	f1 := func() chan struct{} { return nil }
	// 返回只读 channel
	f2 := func() <-chan struct{} { return nil }
	// 返回只写channel
	f3 := func() chan<- struct{} { return nil }
	if false {
		log.Println(&f1, &f2, &f3)
	}
}
