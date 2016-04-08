package gopp

func BytesDup(src []byte) []byte {
	r := make([]byte, len(src))
	n := copy(r, src)
	if n != len(src) {
		panic("wtf")
	}
	return r
}
