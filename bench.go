package gopp

import (
	"log"
	"time"
)

func Benchfn(name string, times int, f func()) {
	btime := time.Now()
	for i := 0; i < times; i++ {
		f()
	}
	dtime := time.Since(btime)
	cps := time.Second / (dtime / time.Duration(times)) // call per second
	log.Println(name, ":\t", times, dtime, dtime/time.Duration(times), int64(cps), "/s")
}
