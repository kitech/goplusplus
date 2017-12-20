package gopp

import (
	"bytes"
	"encoding/json"
	"strings"
	"unicode"

	_ "github.com/huandu/xstrings"
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

// rune support
func Splitrn(s string, n int) []string {
	v := make([]string, 0)

	sub := ""
	sublen := 0
	for _, c := range s {
		cs := string(c)
		if sublen+len(cs) > n {
			v = append(v, sub)
			sub = ""
			sublen = 0
		} else {
			sub += cs
			sublen += len(cs)
		}
	}

	if sublen > 0 {
		v = append(v, sub)
	}
	return v
}

// line support
// TODO one line exceed n???
func Splitln(s string, n int) []string {
	v := make([]string, 0)

	ls := strings.Split(s, "\n")

	sub := ""
	sublen := 0
	for _, line := range ls {
		if sublen+1+len(line) > n {
			v = append(v, sub)
			sub = ""
			sublen = 0
		} else {
			sub += line + "\n"
			sublen += len(line) + 1
		}
	}

	if sublen > 0 {
		v = append(v, sub)
	}
	return v
}

func StrPrepend(s string, b byte) string {
	return string(append([]byte{b}, bytes.NewBufferString(s).Bytes()...))
}

func StrPrepend2(s string, b byte) string {
	return string([]byte{b}) + s
}

func StrReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 仅Title第一个字节
func Title(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(s[:1]) + s[1:]
	}
	return s
}

func IsNumberic(s string) bool {
	if strings.Count(s, ".") > 1 {
		return false
	}
	for _, c := range s {
		if unicode.IsNumber(c) || c == '.' {
		} else {
			return false
		}
	}
	return true
}

func IsInteger(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsPrint(s string) bool {
	for _, c := range s {
		if !unicode.IsPrint(rune(c)) {
			return false
		}
	}
	return true
}

// type String struct{ s string }
/*
func NewString(s string) *String { return &String{s} }
func (this *String) Raw() string { return this.s }

func (this *String) Mid(from, length int) *String { return NewString(this.s[from:length]) }
*/

// 以类方法的方式使用string相关函数，使用时可以拷贝过去
// 不过还是有许多代码要写的

type Str string

func (this Str) Mid(from, length int) Str { return Str(this[from:length]) }

func JsonEncode(v interface{}) (js string, err error) {
	w := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	err = enc.Encode(v)
	js = string(w.Bytes())
	return
}
