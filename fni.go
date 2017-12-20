package gopp

// no use???
// usage: Progn(Retn(f1(...)), Retn(f2(...)), Retn(f3(...)))
func Progn(args ...interface{}) (rets [][]interface{}) {
	for _, arg := range args {
		rets = append(rets, arg.([]interface{}))
	}
	return
}

// 支持>=1个参数的函数
func Retn(args ...interface{}) []interface{} { return args }

// 对于没有返回值的函数，不能作为函数参数传递，所以无法用Progn函数调用。
