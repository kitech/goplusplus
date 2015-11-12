package gopp

import (
	"fmt"
	"reflect"
	"testing"
)

/////
// go test ./...
func TestXChan(t *testing.T) {
	fmt.Println("abc")
	xchan_wait()

	defer xchan_wait()
	xch := NewXChan(2)
	xch.Write(123)
	xch.Write("456")
	xch.PDone()
	xch.PDone()

	v2 := xch.Read()
	fmt.Println(v2, reflect.TypeOf(v2))

	v3 := xch.Read()
	fmt.Println(v3, reflect.TypeOf(v3))

	xch.CDone()
	xch.CDone()
	// xchan_wait()
}

func TestAbc2(t *testing.T) {
	fmt.Println("abc")
	xchan_wait()

}
