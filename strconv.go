package gopp

import "strconv"

func MustInt(s string) int {
	n, err := strconv.Atoi(s)
	ErrPrint(err)
	return n
}
