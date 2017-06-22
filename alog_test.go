package gopp

import (
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestSliceSwapper(t *testing.T) {
	{
		slc := []string{"1", "2"}
		f := reflect.Swapper(slc)
		log.Println(f == nil, &f)
		f(0, 1)
		log.Println(slc) // {"2","1"}
	}
	{
		// f := reflect.Swapper(8) // panic with non slice
		// log.Println(f == nil, f)
	}

	{
		slc := []int{1, 2, 3, 4, 5, 6}
		log.Println(Shuffle2(slc))
		log.Println(Shuffle2(slc))
		log.Println(Shuffle2(slc))
		log.Println(Shuffle2(strings.Split("123456789", "")))
	}
}
