package main

import (
	"fmt"
	"gopp"
	"time"
)

func main() {
	run_thread()

	for {
		select {}
	}
}

func run_thread() {
	f1 := func() {
		time.Sleep(5 * time.Second)
	}

	t := gopp.NewThread(f1)
	fmt.Println(t)
	t.Start()
	// fmt.Println(t.tid) // error
}

func run_pipe_chan() {
	a := gopp.NewXChan(1)
	fmt.Println(a)
}
