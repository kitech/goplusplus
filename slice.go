package gopp

import "log"

// 一个vector/slice中是否包含一个值
func Vcontains(v interface{}, e interface{}) bool {
	ret := false

	fr := func(i interface{}) bool {
		if i == e {
			ret = true
			return true
		}
		return false
	}
	fo := MustFunc(fr)

	res := MapAny(fo, v)
	log.Println(res)

	return ret
}

// 一个 map中是否包含一个值
func Mcontains(m interface{}, e interface{}) bool {
	fr := func(i interface{}, iv interface{}) interface{} {
		return i == e || iv.(bool)
	}
	fo := MustFunc(fr)
	ret := ReduceAny(fo, m, false)
	return ret.(bool)
}

type Vec struct {
	v interface{}
}

func NewVec(v interface{}) *Vec { return &Vec{v} }

type Set struct{}

type Tuple struct {
}

type Triple struct {
}
