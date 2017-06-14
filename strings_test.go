package gopp

import (
	"testing"
)

func TestJsonEncode(t *testing.T) {
	v := make(map[string]string)
	v["v"] = "呵呵"
	v["v2"] = "呵呵ll"

	js, err := JsonEncode(v)
	println(err)
	println(js)
	if err != nil {
		t.Error(err)
	}
}
