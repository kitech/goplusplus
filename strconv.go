package gopp

import (
	"fmt"
	"strconv"
)

func MustInt(s string) int {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	ErrPrint(err, s)
	return n
}

func MustUint32(s string) uint32 {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	ErrPrint(err, s)
	return uint32(n)
}

func MustInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	ErrPrint(err, s)
	return n
}

func ToStr(v interface{}) string { return fmt.Sprintf("%v", v) }
func ToStrs(args ...interface{}) (rets []string) {
	for _, arg := range args {
		rets = append(rets, ToStr(arg))
	}
	return
}
