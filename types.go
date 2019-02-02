package gopp

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

// TODO how add methods for Any type
type IAny interface{}

/*
invalid receiver type *Any (Any is an interface type)
func (this *Any) Hehe() {
}
*/

type Any struct {
	I interface{}
	// v *reflect.Value
}

func ToAny(i interface{}) Any {
	// v := reflect.ValueOf(i)
	return Any{i}
}
func (this Any) Raw() interface{} { return this.I }
func (this Any) IsNil() bool      { return this.I == nil }
func (this Any) I0() int          { return this.I.(int) }
func (this Any) U0() uint         { return this.I.(uint) }
func (this Any) I8() int8         { return this.I.(int8) }
func (this Any) U8() uint8        { return this.I.(uint8) }
func (this Any) I16() int16       { return this.I.(int16) }
func (this Any) U16() uint16      { return this.I.(uint16) }
func (this Any) I32() int32       { return this.I.(int32) }
func (this Any) U32() uint32      { return this.I.(uint32) }
func (this Any) I64() int64       { return this.I.(int64) }
func (this Any) U64() uint64      { return this.I.(uint64) }
func (this Any) F32() float32     { return this.I.(float32) }
func (this Any) F64() float64     { return this.I.(float64) }
func (this Any) Str() string      { return this.I.(string) }
func (this Any) Iterable() bool {
	switch reflect.TypeOf(this.I).Kind() {
	case reflect.Slice, reflect.Array, reflect.Map,
		reflect.String, reflect.Struct:
		return true
	}
	return false
}
func (this Any) CanMKey() bool {
	return true // 是否能作为map的key
}
func (this Any) AsStr() string { return fmt.Sprintf("%v", this.I) }
func (this Any) AsJson() []byte {
	bcc, err := json.Marshal(this.I)
	ErrPrint(err)
	return bcc
}

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

var RefKindTys = map[reflect.Kind]reflect.Type{
	reflect.Int8: Int8Ty, reflect.Uint8: Uint8Ty,
	reflect.Int: IntTy, reflect.Uint: UintTy,
	reflect.Int32: Int32Ty, reflect.Uint32: Uint32Ty,
	reflect.Int64: Int64Ty, reflect.Uint64: Uint64Ty,
	reflect.Float32: Float32Ty, reflect.Float64: Float64Ty,
	reflect.Bool: BoolTy, reflect.String: StrTy,
}

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

const ByteTySz = unsafe.Sizeof(byte(0))
const Int8TySz = strconv.IntSize / 4
const Int16TySz = Int8TySz * 2
const Int32TySz = Int8TySz * 4
const Int64TySz = Int8TySz * 8
const Float64TySz = Int8TySz * 8
const Float32TySz = Int8TySz * 4
const UintptrTySz = unsafe.Sizeof(uintptr(0))
const BoolTySz = unsafe.Sizeof(true)
const IntTySz = unsafe.Sizeof(int(0))
const RuneTySz = unsafe.Sizeof(rune(0))
const EmptyStructTySz = unsafe.Sizeof(struct{}{})

// const NilSz = unsafe.Sizeof(nil)

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
