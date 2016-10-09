package gopp

import (
	"errors"
	"fmt"
	// "log"
	"reflect"
	// "sort"
)

// from functional_go.pdf
type Func struct {
	in  reflect.Type
	out reflect.Type
	f   func(interface{}) interface{}
	f2  func(interface{}, interface{}) interface{}
}

func (self Func) Call(v interface{}) interface{} {
	return self.f(v)
}

func (self Func) Call2(v interface{}, v2 interface{}) interface{} {
	return self.f2(v, v2)
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

// must的含义，断定是某个有效的值，如果不是程序报错退出。
// @param f raw function, like strings.ToUpper
func MustFunc(f interface{}, args ...interface{}) *Func {
	if f == nil {
		panic(args)
	}

	tv := reflect.TypeOf(f)
	if tv.Kind() != reflect.Func {
		panic(args)
	}

	fv, err := NewFunc(f)
	if err != nil || fv == nil {
		panic(err)
	}

	return fv
}

func Must(f *Func, err error) *Func {
	if f == nil {
		panic(err)
	}
	return f
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

// reduce
// func(elem interface{}, iv interface{}) interface{}
func ReduceArrayAny(f *Func, vs []interface{}, iv interface{}) interface{} {
	if len(vs) == 0 {
		return nil
	}
	return ReduceArrayAny(f, vs[:len(vs)-1], f.Call2(vs[len(vs)-1], iv))
}

// func(elem interface{}, iv interface{}) interface{}
func ReduceMapAny(f *Func, vs map[interface{}]interface{}) map[interface{}]interface{} {
	mv := make(map[interface{}]interface{}, 0)
	for k, v := range vs {
		mv[k] = f.Call(v)
	}
	return mv
}

func ReduceAny(f *Func, vs interface{}, iv interface{}) interface{} {
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

func NewMaybe(v interface{}) Maybe {
	return Maybe{v}
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

	res := f.Call(self.Value)
	vr := reflect.ValueOf(res)
	if vr.Kind() == reflect.Ptr && vr.IsNil() {
		return Maybe{}
	}

	return MaybeFrom(res)
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

func NewMany(args ...interface{}) *Many {
	return ManyFrom(args...)
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
		// fmt.Printf("%#v\n", this.Tail)
		res = this.Tail.Map(f)
	}

	v := IfElse(this.Head != nil, f.Call(this.Head), nil)
	v2 := ToSlice(v, true)
	for _, iv := range v2 {
		res = &Many{iv, res}
	}

	return res
}

func (this *Many) Do(fs ...interface{}) *Many {
	if len(fs) == 0 {
		return this
	}

	// f := MustFunc(fs[0])
	f, err := NewFunc(fs[0])

	// 这种应该更合理点吧
	// 相当于忽略所有的执行结果，一个出错，不在给出任何结果，方便执行完成后的检查
	if err != nil {
		return NewMany()
	}

	// 相当于忽略本次执行结果，并继续执行。
	if err != nil {
		// return this
	}

	return this.Map(f).Do(fs[1:]...)
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

///函数组合
func ComposeFunc(f, g *Func) (*Func, error) {
	// tf := reflect.TypeOf(f)
	// tg := reflect.TypeOf(g)

	if f.out != g.in {
		return nil, fmt.Errorf("Can't compose %v != %v\n", f.out, g.in)
	}

	return &Func{
		f.in, g.out,
		func(x interface{}) interface{} { return g.Call(f.Call(x)) },
		// func(x interface{}, y interface{}) interface{} { return g.Call2(f.Call2(x, y), f.Call2(y, x)) },
		nil,
	}, nil
}

func Compose(f, g interface{}) (*Func, error) {
	tf := reflect.TypeOf(f)
	tg := reflect.TypeOf(g)

	if tf.Kind() != reflect.Func || tg.Kind() != reflect.Func {
		return nil, fmt.Errorf("all parameters must be function type.")
	}

	if tf.Out(0) != tg.In(0) {
		return nil, fmt.Errorf("Can't compose %v != %v\n", tf.Out(0), tg.In(0))
	}

	ff, err := NewFunc(f)
	if err != nil {
		return nil, err
	}

	fg, err := NewFunc(g)
	if err != nil {
		return nil, err
	}

	return &Func{
		tf.In(0), tg.Out(0),
		func(x interface{}) interface{} { return fg.Call(ff.Call(x)) },
		// func(x interface{}, y interface{}) interface{} { return fg.Call2(ff.Call2(x, y), ff.Call2(y, x)) },
		nil,
	}, nil
}

///// 常用函数原型
type FilterFunc func(interface{}) bool
type MapFunc func(interface{}) interface{}
type FolderFunc func(interface{}, interface{}) interface{}
type ReduceFunc FolderFunc
