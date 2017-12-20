package gopp

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(b []byte) []byte {
	sum := md5.Sum(b)
	return sum[:]
}

func Md5AsStr(b []byte) string {
	sum := md5.Sum(b)
	return hex.EncodeToString(sum[:])
}

func Md5Str(s string) string { return Md5AsStr([]byte(s)) }
