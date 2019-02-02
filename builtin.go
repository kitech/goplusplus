package gopp

import (
	"encoding/json"
	"log"
	"reflect"
	"unsafe"
)

func BytesReverse(s []byte) []byte {
	r := s
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return r
}

func BytesDup(src []byte) []byte {
	r := make([]byte, len(src))
	n := copy(r, src)
	if n != len(src) {
		panic("wtf")
	}
	return r
}

func DeepCopy(from interface{}, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, to)
}

// deepcopy的一種實現，使用json作爲中轉
// github.com/getlantern/deepcopy
// 還有一種使用reflect遞歸copy所有元素
// https://github.com/mohae/deepcopy

// if integer, number, must use var's addr, like , a := 5; &a
// 一般用于数字类型的操作，指针类型的强制转换
// TODO 考虑类型的存储大小，防止丢失精度
// ptr usage: var s *MyFileStat; OpAssign(&s, &os.FileStat{})
// 返回值为*tovalp
func OpAssign(tovalp, fromvaluep interface{}) interface{} {
	toval := reflect.ValueOf(tovalp)
	log.Println(toval.CanAddr(), toval.Type().String())
	toty := toval.Type()
	fromvalue := reflect.ValueOf(fromvaluep)
	fromty := fromvalue.Type()
	if fromty.Elem().AssignableTo(toty.Elem()) {
		toval.Elem().Set(fromvalue.Elem())
	} else if fromty.Elem().ConvertibleTo(toty.Elem()) {
		toval.Elem().Set(fromvalue.Elem().Convert(toty.Elem()))
	} else if fromty.Kind() == reflect.Ptr && toval.Kind() == reflect.Ptr {
		// 强制指针转换 , (*FileStat)(unsafe.Pointer(*os.fileStat))
		tmp := reflect.NewAt(toval.Type().Elem().Elem(), unsafe.Pointer(fromvalue.Pointer()))
		toval.Elem().Set(tmp)
	} else {
		log.Panicln("Connot assign.", toty.String(), fromty.String())
	}
	return toval.Elem().Interface()
}

func OpEqual(left, right interface{}) bool {
	return false
}

func OpGreatThan(left, right interface{}) bool {
	return false
}
func OpGreatOrEqual(left, right interface{}) bool {
	return false
}
func OpLessThan(left, right interface{}) bool {
	return false
}
func OpLessOrEqual(left, right interface{}) bool {
	return false
}

func Lenv(v interface{}) int {
	switch rv := v.(type) {
	case string:
		return len(rv)
	case byte:
		return int(unsafe.Sizeof(rv))
	case int8:
		return int(unsafe.Sizeof(rv))
	case int:
		return int(unsafe.Sizeof(rv))
	case uint:
		return int(unsafe.Sizeof(rv))
	case int32:
		return int(unsafe.Sizeof(rv))
	case uint32:
		return int(unsafe.Sizeof(rv))
	case int64:
		return int(unsafe.Sizeof(rv))
	case uint64:
		return int(unsafe.Sizeof(rv))
	case float32:
		return int(unsafe.Sizeof(rv))
	case float64:
		return int(unsafe.Sizeof(rv))
	default:
		vv := reflect.ValueOf(v)
		switch vv.Type().Kind() {
		case reflect.Slice:
			return vv.Len()
		case reflect.Array:
			return vv.Len()
		case reflect.Map:
			return vv.Len()
		case reflect.Chan:
			return vv.Len()
		}
	}
	return 0
}
