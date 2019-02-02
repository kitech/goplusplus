package cgopp

/*
#include <stdint.h>
*/
import "C"

import (
	"gopp"
)

func _TestAssign1() bool {
	var to C.uint32_t = 123
	var from int = 567
	gopp.OpAssign(&to, &from)
	return to == C.uint32_t(from)
}
