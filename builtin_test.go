package gopp

import (
	"testing"
	"time"
)

// go test -v -run Channel0 gopp
func TestChannel0(t *testing.T) {
	buf := make([]byte, 8192)
	ch := make(chan []byte, 0)

	times := 1000000
	go func() {
		for i := 0; i < times; i++ {
			ch <- buf
		}
	}()

	btime := time.Now()
	for i := 0; i < times; i++ {
		nbuf := <-ch
		if len(nbuf) != len(buf) {
			t.Fail()
		}
	}
	etime := time.Now()
	dtime := etime.Sub(btime)
	t.Log(etime.Sub(btime), float64(times*len(buf)/1024/1024)/dtime.Seconds(), "MB/s",
		float64(times)/dtime.Seconds())
	// channel数据传输速度很快，10GB/s以上没问题
	// 但是传输次数会怎么样呢
}
