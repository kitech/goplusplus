package gopp

import (
	"os"
	"syscall"
	"time"
)

// try export go internal struct

// A fileStat is the implementation of FileInfo returned by Stat and Lstat.
type FileStat struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	Sys     syscall.Stat_t
}
