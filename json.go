package gopp

import (
	"io"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

type Json struct {
	jo *simplejson.Json
}

func NewEmptyJson() *Json {
	return &Json{simplejson.New()}
}

func NewJson(body []byte) (*Json, error) {
	j, err := simplejson.NewJson(body)
	return &Json{j}, err
}

func NewJsonFromReader(r io.Reader) (*Json, error) {
	j, err := simplejson.NewFromReader(r)
	return &Json{j}, err
}

func NewJsonFromObject(j *simplejson.Json) *Json {
	return &Json{j}
}

func (j *Json) Ori() *simplejson.Json { return j.jo }

func (j *Json) GetPathDot(branch string) *Json {
	return &Json{j.jo.GetPath(strings.Split(branch, ".")...)}
}

/// just wrapper
func (j *Json) GetPath(branch ...string) *Json {
	return &Json{j.jo.GetPath(branch...)}
}

func (j *Json) GetIndex(index int) *Json {
	return &Json{j.jo.GetIndex(index)}
}

func (j *Json) Get(key string) *Json {

	return &Json{j.jo.Get(key)}
}

func (j *Json) CheckGet(key string) (*Json, bool) {
	o, ok := j.jo.CheckGet(key)
	return &Json{o}, ok
}

func (j *Json) CheckGetDot(key string) (*Json, bool) {
	ps := strings.Split(key, ".")

	jo := j.jo
	tok := false
	for i := 0; i < len(ps); i++ {
		tok = false
		if tmpo, ok := jo.CheckGet(ps[i]); ok {
			jo = tmpo
			tok = true
		} else {
			return nil, false
		}
	}

	return &Json{jo}, tok
}

func (j *Json) MustInt(args ...int) int {
	return j.jo.MustInt(args...)
}

func (j *Json) MustString(args ...string) string {
	return j.jo.MustString(args...)
}

func (j *Json) MustArray(args ...[]interface{}) []interface{} {
	return j.jo.MustArray(args...)
}

func (j *Json) MustMap(args ...map[string]interface{}) map[string]interface{} {
	return j.jo.MustMap(args...)
}
