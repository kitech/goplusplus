package gopp

import "sync/atomic"

// TODO general version Picker
// TODO other algo version Picker, like random, consistent hash
type Picker interface {
	SelectOne(items []string, chkfn func(item string) bool) string
}

//
type RRPick struct {
	next uint32
}

func (this *RRPick) SelectOne(items []string, chkfn func(item string) bool) string {
	if len(items) == 0 {
		return ""
	}

	for i := 0; i < len(items); i++ {
		if this.next >= uint32(len(items)) {
			atomic.StoreUint32(&this.next, 0)
		}
		item := items[this.next]
		atomic.AddUint32(&this.next, 1)

		if chkfn != nil {
			if chkfn(item) {
				return item
			}
		} else {
			return item
		}
	}
	return ""
}
