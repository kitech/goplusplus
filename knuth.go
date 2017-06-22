package gopp

import (
	"math/rand"
	"reflect"
)

// rand.Seed from init.go

func Shuffle(slc []interface{}) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}

// 原地重排
func Shuffle2(slc interface{}) interface{} {
	ty := reflect.TypeOf(slc)
	if ty.Kind() != reflect.Slice {
		return nil
	}
	f := reflect.Swapper(slc)
	n := reflect.ValueOf(slc).Len()
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		f(i, j)
	}
	return slc
}

/*
參考：
https://gaohaoyang.github.io/2016/10/16/shuffle-algorithm/
https://gist.github.com/quux00/8258425
*/
