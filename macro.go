package gopp

// FuncAsMacro(funcName, args...)
// funcName必须是在当前源码可引用到的函数，或者是当前包中的函数。
// args 动态适配funcName对应函数的参数表
// 关于返回值的问题，这种实现的宏是可以有返回值的
