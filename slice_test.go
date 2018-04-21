package gopp

import "testing"

func TestVcontains(t *testing.T) {
	v := []string{"abc", "efg", "hik"}
	if !Vcontains(v, "efg") {
		t.Fail()
	}
	if Vcontains(v, "xyz") {
		t.Fail()
	}
}

func TestMcontains(t *testing.T) {
	v := map[int]string{0: "abc", 1: "efg", 2: "hik"}
	if !Mcontains(v, "efg") {
		t.Fail()
	}
	if Mcontains(v, "xyz") {
		t.Fail()
	}
}
