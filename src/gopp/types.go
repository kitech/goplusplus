package gopp

import (
	"reflect"
)

type Any interface{}

// maybe can use Once for lazy
var Int8Ty = reflect.TypeOf(int8(1))
var Uint8Ty = reflect.TypeOf(uint8(1))
var IntTy = reflect.TypeOf(int(1))
var Int32Ty = reflect.TypeOf(int32(1))
var Uint32Ty = reflect.TypeOf(uint32(1))
var Int64Ty = reflect.TypeOf(int64(1))
var Uint64Ty = reflect.TypeOf(uint64(1))
var ByteTy = reflect.TypeOf(byte(1))
var Float32Ty = reflect.TypeOf(float32(1.0))
var Float64Ty = reflect.TypeOf(float64(1.0))
var BoolTy = reflect.TypeOf(true)
var StrTy = reflect.TypeOf("")

func IsMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func IsArray(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Array
}

func IsSlice(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

func IsChan(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Chan
}

func IsFunc(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Func
}

func IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}
