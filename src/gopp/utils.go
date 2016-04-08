package gopp

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"
)

// 简单的三元去处模拟函数
func IfElse(q bool, tv interface{}, fv interface{}) interface{} {
	if q == true {
		return tv
	} else {
		return fv
	}
}

// 把一个值转换为数组切片
// 如果本身即为数组切片，则显式转换为数组类型
// 如果本身不是数组切片，则把该值作为返回数组切片的第一个值。
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

func Assert(v interface{}) {
	if v == nil {
		panic(v)
	}

	tv := reflect.TypeOf(v)
	if tv.Kind() == reflect.Bool && v.(bool) == false {
		panic(v)
	}
}

// 俩工具
// 直接忽略掉变量未使用编译提示
func G_USED(vars ...interface{}) {}
func G_FATAL(err error) {
	if err != nil {
		panic(err)
	}
}

func G_DEBUG(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// 去掉返回值中的error
// 返回值个数不能是变长的
func NOE(v ...interface{}) interface{} {
	n := len(v)
	if n == 0 {
		return nil
	}
	last := v[n-1]
	lt := reflect.TypeOf(last)

	e := errors.New("dummy")

	if lt.Kind() == reflect.TypeOf(e).Kind() {
	}
	return v
}

func WAITIF(condfn func() bool, msec int) {
	for {
		if condfn() {
			break
		}
		time.Sleep(time.Duration(msec) * time.Microsecond)
	}
}

func FileExist(fname string) bool {
	if _, err := os.Stat(fname); err == os.ErrNotExist {
		return false
	}
	return true
}
