package gopp

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	log.SetFlags(log.Flags() | log.LstdFlags | log.Lshortfile)
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

}
