package gopp

import (
	"errors"
	"fmt"
	// "log"
	"reflect"
)

// from functional_go.pdf
type Func struct {
	in  reflect.Type
	out reflect.Type
	f   func(interface{}) interface{}
}

func (self Func) Call(v interface{}) interface{} {
	return self.f(v)
}

func NewFunc(f interface{}) (*Func, error) {
	tf := reflect.TypeOf(f)
	vf := reflect.ValueOf(f)

	if tf.Kind() != reflect.Func {
		return nil, errors.New("not a function")
	}

	if tf.NumIn() > 1 {
		fmt.Println("maybe too much in parameters.")
	}

	if tf.NumOut() > 1 {
		fmt.Println("maybe too much out parameters.")
	}

	return &Func{in: tf.In(0), out: tf.Out(0),
		f: func(x interface{}) interface{} {
			arg_type := reflect.TypeOf(x)
			if arg_type != tf.In(0) {
				if arg_type.ConvertibleTo(tf.In(0)) {
					//fmt.Println("Warning inplict type convertion.")
					arg0 := reflect.ValueOf(x).Convert(tf.In(0))
					res := vf.Call([]reflect.Value{arg0})
					return res[0].Interface()
				} else {
					panic("can not conver from:" + arg_type.Kind().String() + " to " + tf.In(0).Kind().String())
				}
			} else {
				res := vf.Call([]reflect.Value{reflect.ValueOf(x)})
				return res[0].Interface()
			}
		},
	}, nil
}

func Must(f *Func, err error) *Func {
	if f == nil {
		panic(err)
	}
	return f
}

func Must2(v interface{}, args ...interface{}) interface{} {
	if v == nil {
		panic(args)
	}
	return v
}

// func(int) bool
func MapArrayAny(f *Func, vs []interface{}) []interface{} {
	if len(vs) == 0 {
		return nil
	}
	return append(MapArrayAny(f, vs[:len(vs)-1]), f.Call(vs[len(vs)-1]))
}

// func(int) bool
func MapMapAny(f *Func, vs map[interface{}]interface{}) map[interface{}]interface{} {
	mv := make(map[interface{}]interface{}, 0)
	for k, v := range vs {
		mv[k] = f.Call(v)
	}
	return mv
}

func MapAny(f *Func, vs interface{}) interface{} {
	tvs := reflect.TypeOf(vs)
	switch tvs.Kind() {
	case reflect.Map:
		tmp := MapToAnyMap(vs)
		return MapMapAny(f, tmp)
	case reflect.Array:
		tmp := ArrayToAnyArray(vs)
		return MapArrayAny(f, tmp)
		// tmp2 := ToAny(vs)
		// return MapAny(f, tmp2)
	case reflect.String:
		if f.out.Kind() != reflect.Uint8 {
			// panic("must retrun uint8")
		}
		vvs := reflect.ValueOf(vs)
		mstr := ""

		uint8_type := reflect.TypeOf(uint8(1))
		uint8_type = Uint8Ty
		for i := 0; i < vvs.Len(); i++ {
			mv := f.Call(vvs.Index(i).Interface())
			mvt := reflect.TypeOf(mv)
			if mvt != uint8_type {
				mvv := reflect.ValueOf(mv).Convert(uint8_type).Interface()
				mstr += string(mvv.(byte))
			} else {
				mstr += string(mv.(byte))
			}
		}
		return mstr
	case reflect.Slice:
		tmp := reflect.ValueOf(vs)
		res := make([]interface{}, 0)
		for i := 0; i < tmp.Len(); i++ {
			nv := f.Call(tmp.Index(i).Interface())
			res = append(res, nv)
		}
		return res
	default:
		fmt.Println("not impled: " + tvs.Kind().String())
	}
	return nil
}

// tool function
func ToAny(src interface{}) interface{} {
	return src
}

func MapToAnyMap(src interface{}) map[interface{}]interface{} {
	tsrc := reflect.TypeOf(src)
	if tsrc.Kind() != reflect.Map {
		fmt.Println("not a map")
	}

	res := make(map[interface{}]interface{}, 0)
	vsrc := reflect.ValueOf(src)
	for _, k := range vsrc.MapKeys() {
		v := vsrc.MapIndex(k)
		res[k.Interface()] = v.Interface()
	}
	return res
}

func ArrayToAnyMap(src interface{}) map[interface{}]interface{} {
	tsrc := reflect.TypeOf(src)
	if tsrc.Kind() != reflect.Array && tsrc.Kind() != reflect.Slice {
		panic("not a array1:" + tsrc.Kind().String())
	}

	res := make(map[interface{}]interface{}, 0)
	vsrc := reflect.ValueOf(src)
	for i := 0; i < vsrc.Len(); i++ {
		mv := vsrc.Index(i).Interface()
		res[i] = mv
	}
	return res
}

func ArrayToAnyMap2(src interface{}) interface{} {
	return ArrayToAnyMap(src)
}

func ArrayToAnyArray(src interface{}) []interface{} {
	tsrc := reflect.TypeOf(src)
	if tsrc.Kind() != reflect.Array && tsrc.Kind() != reflect.Slice {
		panic("not a array1:" + tsrc.Kind().String())
	}

	res := make([]interface{}, 0)
	vsrc := reflect.ValueOf(src)
	for i := 0; i < vsrc.Len(); i++ {
		v := vsrc.Index(i)
		res = append(res, v.Interface())
	}
	return res
}

func ArrayToAnyArray2(src interface{}) interface{} {
	return ArrayToAnyArray(src)
}

// 特化Map函数示例
func MapArrayInt(f func(v int) bool, vs []int) []bool {
	if len(vs) == 0 {
		return nil
	}

	return append(MapArrayInt(f, vs[:len(vs)-1]), f(vs[len(vs)-1]))
}

/////////
type Maybe struct {
	Value interface{}
}

func MaybeFrom(v interface{}) Maybe {
	return Maybe{v}
}

func (self Maybe) Valid() bool {
	return self.Value == nil
}

func (self Maybe) Map(f *Func) Maybe {
	if self.Value == nil {
		return MaybeFrom(nil)
	}

	return MaybeFrom(f.Call(self.Value))
}

func (self Maybe) Do(fs ...interface{}) Maybe {
	if len(fs) == 0 {
		return self
	}

	f, err := NewFunc(fs[0])
	if err != nil {
		return MaybeFrom(nil)
	}

	return self.Map(f).Do(fs[1:]...)
}

/////////
type Many struct {
	Head interface{}
	Tail *Many
}

func ManyFrom(args ...interface{}) *Many {
	if len(args) == 0 {
		return &Many{}
	}

	if len(args) == 1 {
		return &Many{args[0], ManyFrom()}
	}

	return &Many{args[0], ManyFrom(args[1:]...)}
}

func (this *Many) Map(f *Func) *Many {
	if this == nil {
		return nil
	}

	if this.Head == nil {
		return ManyFrom()
	}

	var res *Many = nil
	if this.Tail != nil {
		fmt.Printf("%#v\n", this.Tail)
		res = this.Tail.Map(f)
	}

	ifelse := func(q bool, tv interface{}, fv interface{}) interface{} {
		if q {
			return tv
		} else {
			return fv
		}
	}

	toSlice := func(v interface{}) []interface{} {
		vt := reflect.TypeOf(v)
		if vt.Kind() == reflect.Slice {
			res := []interface{}{}
			vv := reflect.ValueOf(v)
			for i := 0; i < vv.Len(); i++ {
				res = append(res, vv.Index(i).Interface())
			}
			return res
		} else {
			return []interface{}{v}
		}
	}

	// TODO???
	v := ifelse(this.Head != nil, f.Call(this.Head), nil)
	v2 := toSlice(v)
	for _, iv := range v2 {
		res = &Many{iv, res}
	}

	return res
}

func (this *Many) Count() int {
	c := 0
	if this.Head != nil {
		c += 1
	}

	if this.Tail == nil {
		return c
	}
	return c + this.Tail.Count()
}

func (this *Many) Flat() []interface{} {
	if this.Head == nil {
		return []interface{}{}
	}

	return append([]interface{}{this.Head}, this.Tail.Flat()...)
}
