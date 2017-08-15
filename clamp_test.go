package gopp

import "testing"

func TestClamp0(t *testing.T) {
	if Clamp(1, 2, 3) != 2 {
		t.Fail()
	}
	if Clamp(2, 2, 3) != 2 {
		t.Fail()
	}
	if Clamp(3, 2, 3) != 3 {
		t.Fail()
	}
	if Clamp(4, 2, 3) != 3 {
		t.Fail()
	}
	if Clamp(3, 2, 4) != 3 {
		t.Fail()
	}
}
