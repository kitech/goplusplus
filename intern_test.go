package gopp

import (
	"os"
	"testing"
)

func TestFileStat0(t *testing.T) {
	fi, _ := os.Stat("/tmp/")
	var fio *FileStat
	// fio = (*FileStat)(unsafe.Pointer(val.Pointer()))
	OpAssign(&fio, fi)
	if fio.Name == fi.Name() && fio.ModTime == fi.ModTime() &&
		fio.Size == fi.Size() && fio.Mode == fi.Mode() {
	} else {
		t.Fail()
	}
}
