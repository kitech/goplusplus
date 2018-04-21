package gopp

import (
	"math/rand"
	"time"
)

var r *rand.Rand // Rand for this package.

const LowerChars = "abcdefghijklmnopqrstuvwxyz"
const UpperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const AlphaChars = LowerChars + UpperChars + DigitChars
const DigitChars = "0123456789"
const HexChars = UpperChars + DigitChars
const SymbolChars = "~!@#$%^&*()_+{}|:\"';,./<>?_+"
const PrintableChars = AlphaChars + DigitChars + SymbolChars

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomStringAlphaDigit(strlen int) string {
	const chars = AlphaChars
	return RandomStringAny(strlen, chars)
}

func RandomStringAlphaLower(strlen int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz"
	return RandomStringAny(strlen, chars)
}

func RandomStringAlphaLowerDigit(strlen int) string {
	const chars = LowerChars + DigitChars
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

func RandStrHex(l int) string { return RandomStringAny(l, HexChars) }

// string + string vs []string join的速度
func RandomStringUTF8(strlen int) string {
	var c rune
	r := ""
	for i := 0; i < strlen; i++ {
		c = rand.Int31()
		if len(r+string(c)) > strlen {
			// how?
			// continue // 期待下一个值能够正好凑够字符串长度
		} else if len(r+string(c)) == strlen {
			r += string(c)
			break
		} else {
			r += string(c)
		}
	}
	return r
}

// TODO 首位不能为零
// TODO 小数点呢
func RandomNumber(strlen int) string {
	const chars = DigitChars
	return RandomStringAny(strlen, chars)
}

func RandomDouble(strlen int, faclen int) string {
	const chars = DigitChars
	return RandomStringAny(strlen, chars) +
		"." + RandomStringAny(faclen, chars)
}

func RandomStringAny(strlen int, chars string) string {
	result := ""
	for i := 0; i < strlen; i++ {
		index := r.Intn(len(chars))
		result += chars[index : index+1]
	}
	return result
}
