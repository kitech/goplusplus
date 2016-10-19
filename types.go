package gopp

import (
	"reflect"
)

// TODO how add methods for Any type
type Any interface{}

// maybe can use Once for lazy
var vInt8Ty int8
var vUint8Ty uint8
var vIntTy int
var vUintTy uint
var vInt32Ty int32
var vUint32Ty uint32
var vInt64Ty int64
var vUint64Ty uint64
var vByteTy byte
var vFloat32Ty float32
var vFloat64Ty float64
var vBoolTy bool
var vStrTy string

var Int8Ty = reflect.TypeOf(vInt8Ty)
var Uint8Ty = reflect.TypeOf(vUint8Ty)
var IntTy = reflect.TypeOf(vIntTy)
var UintTy = reflect.TypeOf(vUintTy)
var Int32Ty = reflect.TypeOf(vInt32Ty)
var Uint32Ty = reflect.TypeOf(vUint32Ty)
var Int64Ty = reflect.TypeOf(vInt64Ty)
var Uint64Ty = reflect.TypeOf(vUint64Ty)
var ByteTy = reflect.TypeOf(vByteTy)
var Float32Ty = reflect.TypeOf(vFloat32Ty)
var Float64Ty = reflect.TypeOf(vFloat64Ty)
var BoolTy = reflect.TypeOf(vBoolTy)
var StrTy = reflect.TypeOf(vStrTy)

var Int8PtrTy = reflect.TypeOf(&vInt8Ty)
var Uint8PtrTy = reflect.TypeOf(&vUint8Ty)
var IntPtrTy = reflect.TypeOf(&vIntTy)
var UintPtrTy = reflect.TypeOf(&vUintTy)
var Int32PtrTy = reflect.TypeOf(&vInt32Ty)
var Uint32PtrTy = reflect.TypeOf(&vUint32Ty)
var Int64PtrTy = reflect.TypeOf(&vInt64Ty)
var Uint64PtrTy = reflect.TypeOf(&vUint64Ty)
var BytePtrTy = reflect.TypeOf(&vByteTy)
var Float32PtrTy = reflect.TypeOf(&vFloat32Ty)
var Float64PtrTy = reflect.TypeOf(&vFloat64Ty)
var BoolPtrTy = reflect.TypeOf(&vBoolTy)
var StrPtrTy = reflect.TypeOf(&vStrTy)

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

func IsPtr(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}
