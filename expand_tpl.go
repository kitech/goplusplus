package gopp

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=expand_gen.go gen "expType=BUILTINS"

///// go:generate genny -in=$GOFILE -out=$GOFILE.gen.go gen "KeyType=string,int ValueType=string,int"

type expType generic.Type

func Expand2expType(v []expType) (expType, expType) {
	return v[0], v[1]
}

func Expand3expType(v []expType) (expType, expType, expType) {
	return v[0], v[1], v[2]
}

func Expand4expType(v []expType) (expType, expType, expType, expType) {
	return v[0], v[1], v[2], v[3]
}

func Expand5expType(v []expType) (expType, expType, expType, expType, expType) {
	return v[0], v[1], v[2], v[3], v[4]
}

func Expand6expType(v []expType) (expType, expType, expType, expType, expType, expType) {
	return v[0], v[1], v[2], v[3], v[4], v[5]
}
