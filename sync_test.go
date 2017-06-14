package gopp

import (
	"log"
	"runtime"
	"testing"
	"time"

	deadlock "github.com/sasha-s/go-deadlock"
)

// 读锁也会被写锁阻塞，直到写锁解锁
func TestRWLock(t *testing.T) {
	var mu deadlock.RWMutex
	go func() {
		mu.Lock()
		log.Println("locked w")
		// mu.RLock()
		// log.Println("locked r")
		time.Sleep(3 * time.Second)
		mu.Unlock()
	}()
	runtime.Gosched()

	//*
	go func() {
		log.Println("befer read locked 1")
		mu.RLock()
		log.Println("read locked 1")
		time.Sleep(5 * time.Second)
	}()
	go func() {
		log.Println("befer read locked 2")
		mu.RLock()
		log.Println("read locked 2")
		time.Sleep(5 * time.Second)
	}()
	// */

	time.Sleep(50 * time.Second)
}
