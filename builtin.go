package gopp

func BytesDup(src []byte) []byte {
	r := make([]byte, len(src))
	n := copy(r, src)
	if n != len(src) {
		panic("wtf")
	}
	return r
}

// deepcopy的一種實現，使用json作爲中轉
// github.com/getlantern/deepcopy
// 還有一種使用reflect遞歸copy所有元素
// https://github.com/mohae/deepcopy
