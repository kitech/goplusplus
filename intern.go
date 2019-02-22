package gopp

import (
	"os"
	"time"
)

// try export go internal struct

// A fileStat is the implementation of FileInfo returned by Stat and Lstat.
type FileStat struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	Sys     interface{} //    syscall.Stat_t
}
