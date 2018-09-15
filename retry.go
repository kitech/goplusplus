package gopp

import (
	"fmt"
	"time"
)

const (
	RETRY_GET = iota // get a next retry time
	RETRY_FN
	RETRY_FN_WITH_NO
	RETRY_MODE_MAX
)

const (
	// BO_POWER2
	BO_FIBONACCI = iota
	BO_EXPONENTIAL
	BO_NATRURAL
	BO_FIXED
)

// 1.0
type FixedBackOff struct {
	ntimes int
	name   string
}

func NewFixedBackOff() *FixedBackOff {
	this := &FixedBackOff{}
	this.name = "Fixed"
	return this
}
func (this *FixedBackOff) Next() (int, time.Duration) {
	this.ntimes++
	return this.ntimes, 100 * time.Millisecond
}

// 1.1
type NaturalBackOff struct {
	ntimes int
	name   string
}

func NewNaturalBackOff() *NaturalBackOff {
	this := &NaturalBackOff{}
	this.name = "Natural"
	return this
}

func (this *NaturalBackOff) Next() (int, time.Duration) {
	this.ntimes++
	return this.ntimes, time.Duration(this.ntimes*100) * time.Millisecond
}

// 1.5
type ExponentialBackOff struct {
	initialInterval int     // = 100;//初始间隔
	maxInterval     int     // = 5 * 1000L;//最大间隔
	maxElapsedTime  int     // = 50 * 1000L;//最大时间间隔
	multiplier      float32 //= 1.5;//递增倍数（即下次间隔是上次的多少倍）
	ntimes          int
	name            string
}

func NewExponentialBackOff() *ExponentialBackOff {
	this := &ExponentialBackOff{}
	this.name = "Exponential"
	this.initialInterval = 100
	this.maxInterval = this.initialInterval * 50
	this.maxElapsedTime = this.initialInterval * 500
	this.multiplier = 1.5
	return this
}

func (this *ExponentialBackOff) Next() (int, time.Duration) {
	v := int(float32(this.initialInterval) * this.multiplier)
	this.initialInterval = v
	n := this.ntimes
	this.ntimes++
	return n, time.Duration(v) * time.Millisecond
}

// backoff too quick
// 1.9
type FibonacciBackOff struct {
	no1    int
	no2    int
	ntimes int
	name   string
}

func NewFibonacciBackOff() *FibonacciBackOff {
	this := &FibonacciBackOff{}
	this.name = "Fibonacci"
	this.no1, this.no2 = 0, 100
	return this
}

func (this *FibonacciBackOff) Next() (int, time.Duration) {
	no1, no2 := this.no1, this.no2
	this.no1, this.no2 = no2, no1+no2
	this.ntimes += 1
	return this.ntimes, time.Duration(no1+no2) * time.Millisecond // *100 for uniform unit
}

type RetryBackOff interface{ Next() (int, time.Duration) }

type Retryer struct {
	mode   int
	boff   RetryBackOff
	dofnno func(int) error
	dofn   func() error
}

func (this *Retryer) setBackOff(boty int) {
	switch boty {
	case BO_FIXED:
		this.boff = NewFixedBackOff()
	case BO_NATRURAL:
		this.boff = NewNaturalBackOff()
	case BO_EXPONENTIAL:
		this.boff = NewExponentialBackOff()
	case BO_FIBONACCI:
		fallthrough
	default:
		this.boff = NewFibonacciBackOff()
	}
}

func NewRetry() *Retryer {
	this := &Retryer{}
	this.mode = RETRY_GET
	this.setBackOff(BO_EXPONENTIAL)
	return this
}

func (this *Retryer) NextWait() (ntimes int, nwait time.Duration) {
	return this.boff.Next()
}

func (this *Retryer) NextWaitOnly() time.Duration {
	_, nwait := this.NextWait()
	return nwait
}

///
func NewRetryFn(f func(ntimes int) error) *Retryer {
	this := NewRetry()
	this.mode = RETRY_FN_WITH_NO
	this.dofnno = f
	return this
}

func NewRetryFnOnly(f func() error) *Retryer {
	this := NewRetry()
	this.mode = RETRY_FN
	this.dofn = f
	return this
}

func (this *Retryer) Do(unit time.Duration, ntimes ...int) error {
	return this.do(this.mode == RETRY_FN_WITH_NO, unit, ntimes...)
}

func (this *Retryer) do(withno bool, unit time.Duration, ntimes ...int) (err error) {
	innern := 0
	for {
		if withno {
			err = this.dofnno(innern)
		} else {
			err = this.dofn()
		}
		if err == nil {
			break
		} else {
			if len(ntimes) > 0 && innern > ntimes[0] {
				err = fmt.Errorf("Exceed max ntimes: %d", ntimes)
				break
			} else {
				n, v := this.NextWait()
				innern = n
				time.Sleep(unit * time.Duration(v) / 100) // /100 for uniform unit
			}
		}
	}
	return
}

//
func DoTimes(n int, f func(n int)) {
	for i := 0; i < n; i++ {
		f(i)
	}
}

func DoTimesOnly(n int, f func()) {
	for i := 0; i < n; i++ {
		f()
	}
}
