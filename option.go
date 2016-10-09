package gopp

import (
	"errors"
)

type Option struct {
	v   interface{}
	err error
}

func NewOption(v interface{}, err error) Option {
	return Option{v, err}
}

func Some(v interface{}) Option {
	if v == nil {
		return NewOption(nil, errors.New("Nil value"))
	}
	return NewOption(v, nil)
}

func None() Option {
	return NewOption(nil, errors.New("None"))
}

func (this Option) IsSome() bool {
	return this.err == nil
}

func (this Option) IsNone() bool {
	return this.err != nil
}

func (this Option) And() {

}

func (this Option) AndThen() {

}

func (this Option) Or() {

}

func (this Option) OrElse() {

}

func (this Option) Ok() {

}

func (this Option) OkElse() {

}

func (this Option) To() interface{} {
	return this.v
}

func (this Option) ToInt() int {
	return this.v.(int)
}

func (this Option) ToStr() string {
	return this.v.(string)
}
