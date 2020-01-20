package xlog

/*
#include <stdio.h>
#include <stdarg.h>
*/
import "C"

const skip_depth = 3

func dummy(args ...interface{}) {

}

func Println(args ...interface{}) {
	for i := 0; i < len(args); i++ {

	}
	// dummy(args...) // not work
}

func Fatalln(args ...interface{}) {

}

func Printf(format string, args ...interface{}) {

}

func Keep() {}
