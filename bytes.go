package gopp

import (
	"bytes"
)

func BufAt(buf *bytes.Buffer, pos int) *bytes.Buffer {
	if pos < 0 {
		return nil
	}

	b := buf.Bytes()
	return bytes.NewBuffer(b[pos:])
}

type Buffer struct {
	*bytes.Buffer

	b []byte
}

func NewBufferZero() *Buffer {
	this := &Buffer{}
	this.b = []byte{}
	this.Buffer = bytes.NewBuffer(this.b)
	return this
}

func (this *Buffer) RBufAt(pos int) *Buffer {
	if pos < 0 {
		return nil
	}
	nthis := &Buffer{}
	nthis.b = this.b[pos:]
	nthis.Buffer = bytes.NewBuffer(nthis.b)
	return nthis
}

func (this *Buffer) WBufAt(pos int) *Buffer {
	if pos < 0 {
		return nil
	}
	nthis := &Buffer{}
	this.b = this.Bytes()
	nthis.b = this.b[pos:][0:0] // !!!unbelievable
	nthis.Buffer = bytes.NewBuffer(nthis.b)
	return nthis
}

func NewBufferBuf(b []byte) *Buffer {
	this := &Buffer{}
	this.b = b
	this.Buffer = bytes.NewBuffer(this.b)
	return this
}

func (this *Buffer) Readn(n int) (b []byte, err error) {
	b = make([]byte, n)
	rn, err := this.Read(b)
	b = b[:rn]
	return
}
