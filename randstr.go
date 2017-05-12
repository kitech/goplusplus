package gopp

import (
	"math/rand"
	"time"
)

var r *rand.Rand // Rand for this package.

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomStringAlphaDigit(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	return RandomStringAny(strlen, chars)
}

func RandomStringAlphaLower(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz"
	return RandomStringAny(strlen, chars)
}

func RandomStringAlphaUpper(strlen int) string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return RandomStringAny(strlen, chars)
}

func RandomStringAlphaMixed(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return RandomStringAny(strlen, chars)
}

func RandomStringPrintable(strlen int) string {
	const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+{}|:\"';,./<>?_+"
	return RandomStringAny(strlen, chars)
}

func RandomNumber(strlen int) string {
	const chars = "0123456789"
	return RandomStringAny(strlen, chars)
}

func RandomStringAny(strlen int, chars string) string {
	result := ""
	for i := 0; i < strlen; i++ {
		index := r.Intn(len(chars))
		result += chars[index : index+1]
	}
	return result
}
