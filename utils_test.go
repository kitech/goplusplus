package gopp

import "testing"

func TestNonLazyIfElse(t *testing.T) {
	f1runed := false
	f2runed := false
	f1 := func() int {
		f1runed = true
		return 1
	}
	f2 := func() int {
		f2runed = true
		return 2
	}

	IfElseInt(true, f1(), f2())
	if f1runed && f2runed {
		t.Fail()
	}
}

type LazyAny chan func() interface{}

func TestLazyIfElse(t *testing.T) {
	f1runed := false
	f2runed := false
	f1 := func() int {
		f1runed = true
		return 1
	}
	f2 := func() int {
		f2runed = true
		return 2
	}

	IfElseInt(true, Any{f1()}.I0(), Any{f2()}.I0())
	if f1runed && f2runed {
		t.Log(f1runed, f2runed)
		t.Fail()
	}
}
