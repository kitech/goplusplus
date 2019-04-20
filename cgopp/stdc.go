package cgopp

/*
#include <string.h>
#include <stdlib.h>
*/
import "C"

import "gopp"

// std c library functions
// 这么封装一次，引用的包不需要再显式的引入"C"包。让CGO代码转换不传播到引用的代码中
func CMemcpy()  {}
func CFree()    {}
func Calloc()   {}
func CMemset()  {}
func CMemZero() {}
func CStrcpy()  {}
func CStrdup()  {}

const CBoolTySz = gopp.Int32TySz
const CppBoolTySz = gopp.Int8TySz
