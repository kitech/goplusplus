package gopp

import (
	"context"
	"testing"
	"time"
)

func TestCtx1(t *testing.T) {
	c1 := context.Background()
	c2, ccfn1 := context.WithCancel(c1)
	c3, ccfn2 := context.WithDeadline(c1, time.Now().Add(5*time.Second))
	c4, ccfn3 := context.WithTimeout(c1, 5*time.Second)
	c5 := context.WithValue(c1, "id", "87703232")

	G_UNUSED(c1, c2, c3, c4, c5, &ccfn1, &ccfn2, &ccfn3)

}
