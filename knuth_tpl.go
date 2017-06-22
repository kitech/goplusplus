package gopp

import "math/rand"

//go:generate genny -in=$GOFILE -out=knuth_gen.go gen "expType=BUILTINS"

func ShuffleexpType(slc []expType) {
	n := len(slc)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slc[i], slc[j] = slc[j], slc[i]
	}
}
