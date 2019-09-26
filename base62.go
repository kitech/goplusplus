package gopp

import (
	"encoding/base64"
	"fmt"

	"github.com/lytics/base62"
)

// github.com/lytics/base62

const b62encs = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-0123456789."

var b62enc = base62.NewEncoding(b62encs)

func Base62Encode(s string) string {
	return b62enc.EncodeToString([]byte(s))
}

func Base62Decode(s string) (b []byte, err error) {
	defer func() {
		if errx := recover(); errx != nil {
			err = fmt.Errorf("%v", recover())
		}
	}()
	return b62enc.DecodeString(s)
}
func Base62DecStr(s string) (string, error) {
	b, err := Base62Decode(s)
	return string(b), err
}

// url safe: + => -, / => _
const encodeURL = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"

var rawb64enc = base64.NewEncoding(encodeURL).WithPadding(base64.NoPadding)

func Base64EncSafe(s string) string {
	return rawb64enc.EncodeToString([]byte(s))
}
func Base64DecSafe(s string) (b []byte, err error) {
	defer func() {
		if errx := recover(); errx != nil {
			err = fmt.Errorf("%v", recover())
		}
	}()
	b, err = rawb64enc.DecodeString(s)
	return
}
func Base64DecSafeStr(s string) (r string, err error) {
	b, err := Base64DecSafe(s)
	return string(b), err
}

func Encb64Std(s string) string    { return base64.StdEncoding.EncodeToString([]byte(s)) }
func Base64EncStd(s string) string { return base64.StdEncoding.EncodeToString([]byte(s)) }
func Base64DecStd(s string) (b []byte, err error) {
	return base64.StdEncoding.DecodeString(s)
}
func Decb64Std(s string) (b []byte, err error) {
	return base64.StdEncoding.DecodeString(s)
}

func Decb64StdStr(s string) (r string, err error) { return Base64DecStdStr(s) }
func Base64DecStdStr(s string) (r string, err error) {
	b, err := base64.StdEncoding.DecodeString(s)
	return string(b), err
}
