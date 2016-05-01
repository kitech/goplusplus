package gopp

import (
	"bytes"
	"strings"
)

// 安全提取子字符串。支持负值，表示从后面
func SubStr(s string, n int) string {
	if n < 0 {
		absn := AbsI32(n)
		if absn > len(s) {
			return s
		}
		return s[len(s)-absn:]
	} else {
		if n >= len(s) {
			return s
		}
		return s[:n]
	}
}

func StrSuf(s string, n int) string {
	r := SubStr(s, n)
	if len(r) < len(s) {
		return r + "..."
	}
	return r
}

func SubBytes(p []byte, n int) []byte {
	if n >= len(p) {
		return p
	}
	return p[:n]
}

// 按长度切割字符串
func Splitn(s string, n int) []string {
	v := make([]string, 0)
	for i := 0; i < (len(s)/n)+1; i++ {
		bp := i * n
		ep := bp + n
		if bp >= len(s) {
			break
		}
		if ep > len(s) {
			ep = len(s)
		}

		v = append(v, s[bp:ep])
	}
	return v
}

func StrPrepend(s string, b byte) string {
	return string(append([]byte{b}, bytes.NewBufferString(s).Bytes()...))
}

func StrPrepend2(s string, b byte) string {
	return string([]byte{b}) + s
}

// 仅Title第一个字节
func Title(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}
