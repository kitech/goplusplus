package gopp

import (
	"log"
	"reflect"
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

// 測試一下chan的運行時狀態
func TestChanRT0(t *testing.T) {
	{ // close不影響channel 的valid屬性
		c := make(chan bool, 0)
		cv := reflect.ValueOf(c)
		log.Println(cv.IsValid()) //true
		close(c)
		log.Println(cv.IsValid()) //true
	}
	{ // close的channel讀取到默認值
		c := make(chan bool, 0)
		cv := reflect.ValueOf(c)
		x, ok := cv.TryRecv()
		log.Println(x, ok) // <invalid reflect.Value> false
		close(c)
		x, ok = cv.TryRecv()
		log.Println(x, ok) // false,false
	}
	{ // close 的 channel 寫入導致panic
		c := make(chan bool, 0)
		cv := reflect.ValueOf(c)
		ok := cv.TrySend(reflect.ValueOf(true))
		log.Println(ok) // false
		close(c)
		// ok = cv.TrySend(reflect.ValueOf(true))
		// log.Println(ok) // panic
		err := SafeTrySend(c, true)
		log.Println(err) //  send on closed channel
	}
}
