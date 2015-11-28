package gopp

import (
	"reflect"
)

func IfElse(q bool, tv interface{}, fv interface{}) interface{} {
	if q {
		return tv
	} else {
		return fv
	}
}

func ToSlice(v interface{}, reverse bool) []interface{} {
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Slice {
		res := []interface{}{}
		vv := reflect.ValueOf(v)
		for i := 0; i < vv.Len(); i++ {
			idx := IfElse(reverse, vv.Len()-i-1, i).(int)
			res = append(res, vv.Index(idx).Interface())
		}
		return res
	} else {
		return []interface{}{v}
	}
}
