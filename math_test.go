package gopp

import (
	"log"
	"testing"
)

func TestAbsNum0(t *testing.T) {
	xa := AbsNum(-5)
	log.Println(xa)
	xa = AbsNum(int32(-6))
	log.Println(xa)
	xa = AbsNum(float32(-7.5))
	log.Println(xa)
}
