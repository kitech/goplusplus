// error with stack
package gopp

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"runtime/debug"

	_ "github.com/pkg/errors"
)

// Error with errno and stack info
type Error struct {
	errno  int
	errstr string
	stack  []uintptr
}

// eno 模擬可選參數
func NewError(estr string, eno ...int) Error {
	return newErrorN(estr, 2, eno...)
}
func NewErrora(ev interface{}, eno ...int) Error {
	return newErrorN(fmt.Sprintf("%v", ev), 2, eno...)
}
func NewErroraN(ev interface{}, skipn int, eno ...int) Error {
	return newErrorN(fmt.Sprintf("%v", ev), skipn+2, eno...)
}

func newErrorN(estr string, skipn int, eno ...int) Error {
	var pc = make([]uintptr, 0)
	n := runtime.Callers(skipn, nil)
	pc = make([]uintptr, n)
	runtime.Callers(skipn, pc)
	if eno != nil && len(eno) > 0 {
		return Error{errno: eno[0], errstr: estr, stack: pc}
	}
	return Error{errno: 0, errstr: estr, stack: pc}
}

func ErrorFrom(e error) Error {
	return NewError(e.Error())
}

func (this Error) Errno() int {
	return this.errno
}

func (this Error) Errstr() string {
	return this.errstr
}

func (this Error) Error() string {
	return this.String()
}

func (this Error) String() string {
	return fmt.Sprintf("Error: %d, %s", this.errno, this.errstr)
}

func (this Error) PrintStack() {
	fmt.Println(this.String())
	fmt.Println()
	for idx, pc := range this.stack {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		fmt.Printf("#%d, %s %s:%d\n", idx, fn.Name(), file, line)
	}
}

func (this Error) Show() {
	this.PrintStack()
}

func (this Error) Display() {
	this.PrintStack()
}

func printq(v interface{}, args ...interface{}) string {
	msg := fmt.Sprintf("%+v", v)
	for _, arg := range args {
		msg += fmt.Sprintf(" %+v", arg)
	}
	return msg
}

func ErrPrint(err error, args ...interface{}) error {
	if err != nil {
		log.Output(2, printq(err, args...))
	}
	return err
}

func ErrFatal(err error) {
	if err != nil {
		log.Output(2, fmt.Sprintf("%v", err))
	}
}

func FalsePrint(ok bool, args ...interface{}) bool {
	if !ok {
		log.Output(2, printq(ok, args...))
	}
	return ok
}

func NilPrint(v interface{}, args ...interface{}) interface{} {
	if v == nil {
		log.Output(2, printq(v, args...))
	}
	return v
}

func ZeroPrint(v interface{}, args ...interface{}) interface{} {
	if reflect.Zero(reflect.TypeOf(v)).Interface() == v {
		log.Output(2, printq(v, args))
	}
	return v
}

// NOT mean: error != nil or bool == false or int == 0 or pointer == nil or other what?
func NotPrint(v interface{}, args ...interface{}) interface{} {
	switch rv := v.(type) {
	case error:
		ErrPrint(rv, args...)
	case bool:
		FalsePrint(rv, args...)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		ZeroPrint(rv, args...)
	default:
		vty := reflect.TypeOf(v)
		switch vty.Kind() {
		case reflect.Ptr, reflect.Func, reflect.Interface,
			reflect.Array, reflect.Chan, reflect.Map, reflect.Slice,
			reflect.UnsafePointer:
			NilPrint(v, args...)
		default:
			log.Println("unknown type:", vty.String())
		}
	}
	return v
}

func init() {
	if false {
		f1 := func() error {
			return NewError("hehe")
		}
		if f1 != nil {
		}
	}
}

// usage: defer Panicp()
func Panicp() {
	if err := recover(); err != nil {
		bs := debug.Stack()
		log.Println("error:", err, ", stack:", string(bs))
	}
}

// usage: defer func(){Panicp(recover())}()
// do use like this: defer gopp.Panicp(recover())
// because need lazy call recover()
// need macro in golang
func Panicp2(err interface{}) {
	if err != nil {
		bs := debug.Stack()
		log.Println("error:", err, ", stack:", string(bs))
	}
}
