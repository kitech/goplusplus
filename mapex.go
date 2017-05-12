package gopp

import (
	"reflect"
)

// map的更多的操作

func MapKeys(m interface{}) []interface{} {
	mt := reflect.ValueOf(m)
	if mt.Kind() != reflect.Map {
		return nil
	}

	mv := reflect.ValueOf(m)
	return Value2Iface(mv.MapKeys())
}

func MapValues(m interface{}) []interface{} {
	mt := reflect.ValueOf(m)
	if mt.Kind() != reflect.Map {
		return nil
	}

	outs := make([]interface{}, 0)
	mv := reflect.ValueOf(m)
	for _, vk := range mv.MapKeys() {
		vv := mv.MapIndex(vk).Interface()
		outs = append(outs, vv)
	}
	return outs
}

type Map struct {
}

type Array struct {
	a Any
}

func NewArray(a interface{}) *Array {
	this := &Array{}
	this.a = ToAny(a)
	return this
}

func (this *Array) Contains(i interface{}) bool {
	return false
}

func (this *Array) Length() int {
	return 0
}
