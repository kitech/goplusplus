package gopp

import (
	"reflect"
)

type Pair struct {
	Key   interface{}
	Val   interface{}
	Extra interface{}
}

func NewPair(key, val interface{}, extra ...interface{}) *Pair {
	p := &Pair{key, val, nil}
	if len(extra) > 0 {
		p.Extra = extra[0]
	}
	return p
}

// 支持可以迭代的类型：结构体，slice，数组，字符串，map
func Domap(ins interface{}, f func(interface{}) interface{}) (outs []interface{}) {
	outs = make([]interface{}, 0)

	tmpty := reflect.TypeOf(ins)
	// the same as DomapTypeField
	if tmpty.Kind() == reflect.Ptr && tmpty.String() == "*reflect.rtype" {
		insty := ins.(reflect.Type)

		for idx := 0; idx < insty.NumField(); idx++ {
			field := insty.Field(idx)
			out := f(field)
			outs = append(outs, out)
		}
	} else if tmpty.Kind() == reflect.Slice || tmpty.Kind() == reflect.Array {
		tmpv := reflect.ValueOf(ins)
		for i := 0; i < tmpv.Len(); i++ {
			out := f(tmpv.Index(i).Interface())
			outs = append(outs, out)
		}
	} else if tmpty.Kind() == reflect.Map {
		tmpv := reflect.ValueOf(ins)
		for _, vk := range tmpv.MapKeys() {
			out := f(tmpv.MapIndex(vk).Interface())
			outs = append(outs, &Pair{vk.Interface(), out, nil})
		}
	} else {
		insRanger := ins.([]interface{})
		for _, in := range insRanger {
			out := f(in)
			outs = append(outs, out)
		}
	}

	return
}

func DomapTypeField(ty reflect.Type, f func(reflect.StructField) interface{}) (outs []interface{}) {
	outs = make([]interface{}, 0)

	for idx := 0; idx < ty.NumField(); idx++ {
		field := ty.Field(idx)
		out := f(field)
		outs = append(outs, out)
	}

	return
}

func Doreduce(ins interface{}, v interface{},
	f func(v, i interface{}) interface{}) interface{} {
	tmpty := reflect.TypeOf(ins)
	// the same as DomapTypeField
	if tmpty.Kind() == reflect.Ptr && tmpty.String() == "*reflect.rtype" {
		insty := ins.(reflect.Type)

		for idx := 0; idx < insty.NumField(); idx++ {
			field := insty.Field(idx)
			v = f(v, field)
		}
	} else if tmpty.Kind() == reflect.Slice || tmpty.Kind() == reflect.Array {
		tmpv := reflect.ValueOf(ins)
		for i := 0; i < tmpv.Len(); i++ {
			v = f(v, tmpv.Index(i).Interface())
		}
	} else if tmpty.Kind() == reflect.Map {
		tmpv := reflect.ValueOf(ins)
		for _, vk := range tmpv.MapKeys() {
			v = f(v, tmpv.MapIndex(vk).Interface())
		}
	} else {
		insRanger := ins.([]interface{})
		for _, in := range insRanger {
			v = f(v, in)
		}
	}

	return v
}

// interface vector to strings
func IV2Strings(items []interface{}) []string {
	if items == nil {
		return nil
	}

	rets := make([]string, 0)
	for idx := 0; idx < len(items); idx++ {
		rets = append(rets, items[idx].(string))
	}
	return rets
}
func Strs2IV(items []string) []interface{} {
	if items == nil {
		return nil
	}

	rets := make([]interface{}, 0)
	for idx := 0; idx < len(items); idx++ {
		rets = append(rets, items[idx])
	}
	return rets
}

// enumerate类似功能
// 第一种方式，采用数组,可能用内存比较多
// usage: for i := range gopp.Range(5){}
func RangeA(n int) (rg []int) {
	rg = make([]int, n)
	for i := 0; i < n; i++ {
		rg[i] = i
	}
	return
}

// 第二种方式，采用channel。由于用到一个goroutine，可能效率慢
func RangeC(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

// TODO
type Iterable interface {
	Iter() interface{}
	Next() interface{}
}

// string/map/slice/struct or implementation Iterable
func CanIter(v interface{}) bool {
	return false
}
