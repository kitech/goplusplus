package gopp

// Inspired by ramda.js and github.com/azbshiri/ramda

type Ramda interface {
	Always(a interface{}) func() interface{}
	And(a, b bool) bool
	Equals(a ...interface{}) func() interface{}
	Gt(a, b interface{}) bool
	Gte(a, b interface{}) bool
	Head(s interface{}) interface{}
	Identity(a interface{}) interface{}
	Inc(a interface{}) interface{}
	// NewCond(pairs Cond) func(interface{}) interface{}
	T() func() interface{}
	Tail(s interface{}) interface{}
}
