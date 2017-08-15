package gopp

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/emirpasic/gods/sets/hashset"
)

func KeepMe() {}

// 直接传递候选函数的方式
/*
类似C++的overload name resolve

@param args 传递的实参值
@param fns 候选函数列表
@return 返回候选函数索引值，失败返回-1
*/
func SymbolResolveFns(args []interface{}, fns []interface{}) int {
	vtys := make(map[int]map[int]reflect.Type)

	for idx, fnx := range fns {
		vty := getFuncTypes(fnx)
		vtys[idx] = vty
	}

	return SymbolResolve(args, vtys)
}

// 类似C++的overload name resolve
/*
@param args 传递的实参值
@param vtys 候选函数形参类型
@return 返回候选函数索引值，失败返回-1
*/
func SymbolResolve(args []interface{}, vtys map[int]map[int]reflect.Type) int {
	// cycle 1
	if ret := symbolResolveComplete(args, vtys); ret != -1 {
		return ret
	}

	// cycle 2
	if ret := symbolResolveSubTypeConvert(args, vtys); ret != -1 {
		return ret
	}

	// cycle 3
	if ret := symbolResolveConvert(args, vtys); ret != -1 {
		return ret
	}

	// cycle others
	if ret := symbolResolveFind(args, vtys, canHandyConvertFinder); ret != 0-1 {
		return ret
	}

	return -1
}

func FillDefaultValues(args []interface{}, dvals []interface{}) []interface{} {
	return args
}

// 完全匹配，参数个数，参数类型
func symbolResolveComplete(args []interface{}, vtys map[int]map[int]reflect.Type) int {
	return symbolResolveFind(args, vtys, assignFinder)
}

// 参数个数匹配
func symbolResolveCount() {

}

// 参数类型直接匹配？是否是symbolResolveCompelete？
func symbolResolveType() {
}

func symbolResolveSubTypeConvert(args []interface{}, vtys map[int]map[int]reflect.Type) int {
	return symbolResolveFind(args, vtys, convertSubTypeFinder)
}

// TODO convert区分子类型与优先级, 比如float32类型匹配float64类型优先
// 子类型，float: float32/float64
// 参数个数匹配
// 参数类型转换匹配
func symbolResolveConvert(args []interface{}, vtys map[int]map[int]reflect.Type) int {
	return symbolResolveFind(args, vtys, convertFinder)
}

type symfindfunc func(idx int, fromty, toty reflect.Type) bool

func assignFinder(idx int, fromty, toty reflect.Type) bool {
	return fromty.AssignableTo(toty)
}

func convertFinder(idx int, fromty, toty reflect.Type) bool {
	return fromty.ConvertibleTo(toty)
}

func convertSubTypeFinder(idx int, fromty, toty reflect.Type) bool {
	if fromty.ConvertibleTo(toty) {
		subsetsF := hashset.New()
		subsetsF.Add(reflect.Float32, reflect.Float64)
		if subsetsF.Contains(fromty.Kind(), toty.Kind()) {
			return true
		}
		subsetsI := hashset.New()
		subsetsI.Add(reflect.Int, reflect.Int8,
			reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64)
		if subsetsI.Contains(fromty.Kind(), toty.Kind()) {
			return true
		}
	}
	return false
}

func canHandyConvertFinder(idx int, fromty, toty reflect.Type) bool {
	if toty.Kind() != reflect.Struct &&
		(fromty.ConvertibleTo(toty) || fromty.AssignableTo(toty)) {
		return true
	}

	if canHandyConvert(fromty, toty) {
		return true
	}

	return false
}

// TODO FindAll，没有全局信息，无法找到最优匹配
// 查找满足条件的第一条
func symbolResolveFind(args []interface{}, vtys map[int]map[int]reflect.Type, matcher symfindfunc) int {
	argc := len(args)
	if argc < 0 {
	}

	matsyms := make(map[int]bool, 0)
	for symidx := 0; symidx < len(vtys); symidx++ {
		vty := vtys[symidx]
		pnum := len(vty)

		if argc != pnum {
			continue
		}

		// 参数类型匹配
		matp := make(map[int]bool, len(vty))
		for idx := 0; idx < len(vty); idx++ {
			matp[idx] = false

			ety := vty[idx]
			aty := reflect.TypeOf(args[idx])

			// 关键点？
			if matcher(idx, aty, ety) {
				matp[idx] = true
				if ety.Size() < aty.Size() {
					log.Println("Warning: maybe lost precision")
				}
			}
		}

		// a folder/reduce
		ismat := true
		for idx := 0; idx < len(matp); idx++ {
			ismat = ismat && matp[idx]
		}

		// 查找到第一个match的symbol即返回
		if ismat {
			matsyms[symidx] = true // 准备多级match
			return symidx
		}
	}

	return -1
}

// 无临时对象的参数类型转换匹配
func symbolResolveNotemp() {

}

func canHandyConvert(from reflect.Type, to reflect.Type) bool {
	infos := make([]string, 0)

	infos = append(infos, from.Kind().String())
	switch from.Kind() {
	case reflect.Ptr:
		infos = append(infos, from.Elem().Kind().String(), from.Elem().Name())
	case reflect.Struct:
		infos = append(infos, from.Name())
	default:
	}

	infos = append(infos, "===")

	infos = append(infos, to.Kind().String())
	switch to.Kind() {
	case reflect.Ptr:
		infos = append(infos, to.Elem().Kind().String(), to.Elem().Name())
	case reflect.Struct:
		infos = append(infos, to.Name())
	default:
	}

	log.Println(strings.Join(infos, ", "))

	//
	switch {
	// string => char *
	case from.Kind() == reflect.String &&
		to.Kind() == reflect.Ptr && to.Elem().Kind() == reflect.Uint8:
		return true
		// &QXxx => QXxx
	case from.Kind() == reflect.Ptr &&
		from.Elem().Kind() == to.Kind() && from.Elem().Name() == to.Name():
		return true
		// QXxx => QXxx
	case from.Kind() == reflect.Struct &&
		from.Kind() == to.Kind() && from.Name() == to.Name():
		return true
	}
	return false
}

func ErrorResolve(class, method string, args []interface{}) {
	pcs := make([]uintptr, 3)
	npc := runtime.Callers(0, pcs)
	if npc == -1 {
	}
	rtf := runtime.FuncForPC(pcs[2])
	file, line := rtf.FileLine(pcs[2])
	fmt.Println(rtf.Name(), file, line, "Unresolved VT", class, method, args)
}

/////////
// runtime.SetFinalizer(x, UniverseFree)
func UniverseFree(this interface{}) {
	oty := reflect.TypeOf(this)
	oval := reflect.ValueOf(this)

	_, ok := oty.MethodByName("Free")
	if ok {
		// in := []reflect.Value{oval}
		// mth.Func.Call(in)
		oval.MethodByName("Free").Call([]reflect.Value{})
		// log.Println(this, "freed", oty.Elem().Name())
	} else {
		log.Println(this, "has no Free method.", oty.Elem().Name())
	}
}

/////////
func getFuncTypes(f interface{}) map[int]reflect.Type {
	if f == nil {
		return nil
	}

	vf := reflect.TypeOf(f)
	if vf.Kind() != reflect.Func {
		log.Fatalln(vf)
		return nil
	}

	rets := make(map[int]reflect.Type, 0)
	for idx := 0; idx < vf.NumIn(); idx++ {
		rets[idx] = vf.In(idx)
	}
	return rets
}

func getArgsValues(args []interface{}) (argv []reflect.Value) {
	argv = make([]reflect.Value, len(args))
	for idx, arg := range args {
		argv[idx] = reflect.ValueOf(arg)
	}
	return
}

/////////

func StringTy(pointer bool) reflect.Type {
	var s = "foo"
	if pointer {
		return reflect.TypeOf(&s)
	}
	return reflect.TypeOf(s)
}

func RuneTy(pointer bool) reflect.Type {
	var s rune = '\000'
	if pointer {
		return reflect.TypeOf(&s)
	}
	return reflect.TypeOf(s)
}

func Int16Ty(pointer bool) reflect.Type {
	if pointer {
		var v = int16(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(int16(0))
}

func UInt16Ty(pointer bool) reflect.Type {
	if pointer {
		var v = uint16(0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(uint16(0))
}

func FloatTy(pointer bool) reflect.Type {
	if pointer {
		var v = float32(0.0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(float32(0.0))
}
func DoubleTy(pointer bool) reflect.Type {
	if pointer {
		var v = float64(0.0)
		return reflect.TypeOf(&v)
	}
	return reflect.TypeOf(float64(0.0))
}

func VoidpTy() reflect.Type {
	var v unsafe.Pointer = nil
	return reflect.TypeOf(v)
}

/// overload helper function
func convArgsForFunc(args []interface{}, fn interface{}) (argv []reflect.Value) {
	argv = make([]reflect.Value, len(args))

	fnty := reflect.TypeOf(fn)
	for i := 0; i < fnty.NumIn(); i++ {
		argv[i] = convArgForType(args[i], fnty.In(i))
	}

	return
}

func convArgForType(arg interface{}, ty reflect.Type) reflect.Value {
	argval := reflect.ValueOf(arg)
	retval := argval.Convert(ty)
	return retval
}

func overloadRcCall(args []interface{}, fnx interface{}) (out []reflect.Value) {
	in := convArgsForFunc(args, fnx)
	out = reflect.ValueOf(fnx).Call(in)
	return out
}
