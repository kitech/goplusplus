package gopp

import (
	"log"
)

func init() {
	log.SetFlags(log.Flags() | log.LstdFlags | log.Lshortfile)
}
