package gopp

import "strconv"

func MustInt(s string) int {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	ErrPrint(err)
	return n
}

func MustUint32(s string) uint32 {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	ErrPrint(err)
	return uint32(n)
}
