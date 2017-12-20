package gopp

import "github.com/lytics/base62"

// github.com/lytics/base62

func Base62Encode(s string) string {
	return base62.StdEncoding.EncodeToString([]byte(s))
}

func Base62Decode(s string) ([]byte, error) {
	return base62.StdEncoding.DecodeString(s)
}
