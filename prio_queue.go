package gopp

import (
	"container/heap"
	"math"
	"sync"
)

// weight大，优先级高
type PQItem interface {
	Key() []interface{}
	Weight() int
	// Less(that PQItem) bool
}

// mixed queue and map features
type PrioQueue struct {
	mu     sync.RWMutex
	itemsv priovec
	itemsm map[interface{}]*pqitemin // key =>
	maxcnt int                       // TODO 并未实现
}

func NewPrioQueue(maxcnt ...int) *PrioQueue {
	q := &PrioQueue{}
	q.itemsv = priovec{}
	q.itemsm = map[interface{}]*pqitemin{}
	heap.Init(&q.itemsv)
	if len(maxcnt) > 0 {
		q.maxcnt = maxcnt[0]
	} else {
		q.maxcnt = math.MinInt32
	}
	return q
}

// 仅push不存在的, PushNew
func (pq *PrioQueue) Push(item PQItem) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	exist := false
	keys := item.Key()
	for _, key := range keys {
		olditem, ok := pq.itemsm[key]
		if ok {
			_ = olditem
			// olditem.value = item
			// heap.Fix(&pq.itemsv, olditem.index)
		}
		exist = exist || ok
	}
	if !exist {
		itemin := &pqitemin{value: item}
		heap.Push(&pq.itemsv, itemin)
		for _, key := range keys {
			pq.itemsm[key] = itemin
		}
	}
	pq.keepcount(item)
}
func (pq *PrioQueue) PushSet(item PQItem) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	exist := false
	var olditem1 *pqitemin
	keys := item.Key()
	for _, key := range keys {
		olditem, ok := pq.itemsm[key]
		if ok {
			_ = olditem
			olditem1 = olditem
			// olditem.value = item
			// heap.Fix(&pq.itemsv, olditem.index)
		}
		exist = exist || ok
		if exist {
			break
		}
	}
	if exist {
		olditem1.value = item
		heap.Fix(&pq.itemsv, olditem1.index)
	} else {
		itemin := &pqitemin{value: item}
		heap.Push(&pq.itemsv, itemin)
		for _, key := range keys {
			pq.itemsm[key] = itemin
		}
	}
	pq.keepcount(item)
}
func (pq *PrioQueue) keepcount(item PQItem) {
	// lock by Push/PushSet
	if pq.maxcnt <= 0 {
		return
	}
	curlen := len(pq.itemsv)
	if curlen <= pq.maxcnt {
		return
	}
	var items []interface{}
	for i := 0; i < pq.maxcnt; i++ {
		items = append(items, heap.Pop(&pq.itemsv))
	}
	for pq.itemsv.Len() > 0 {
		itemx := heap.Pop(&pq.itemsv)
		item := itemx.(*pqitemin)
		for _, key := range item.value.Key() {
			delete(pq.itemsm, key)
		}
		// log.Println("deled", item.value.Key(), item.value.Weight())
	}

	for _, item := range items {
		heap.Push(&pq.itemsv, item)
	}
	// log.Println("added", item.Key(), item.Weight())
}

func (pq *PrioQueue) Pop() PQItem {
	pq.mu.Lock()
	defer pq.mu.Unlock()
	if len(pq.itemsv) == 0 {
		return nil
	}
	itemx := heap.Pop(&pq.itemsv)
	item := itemx.(*pqitemin)
	keys := item.value.Key()
	for _, key := range keys {
		delete(pq.itemsm, key)
	}
	return item.value
}

func (pq *PrioQueue) Len() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	// Assert(len(pq.itemsv) == len(pq.itemsm), "Ooops, invalid inner state", len(pq.itemsv), len(pq.itemsm))
	return len(pq.itemsv)
}
func (pq *PrioQueue) Cap() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	// Assert(len(pq.itemsv) == len(pq.itemsm), "Ooops, invalid inner state", len(pq.itemsv), len(pq.itemsm))
	return pq.maxcnt
}

func (pq *PrioQueue) Del(key interface{}) PQItem {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	item, ok := pq.itemsm[key]
	if ok {
		heap.Remove(&pq.itemsv, item.index)
		keys := item.value.Key()
		for _, key := range keys {
			delete(pq.itemsm, key)
		}
		return item.value
	}
	return nil
}

// readonly
func (pq *PrioQueue) GetKey(key interface{}) PQItem {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	item, ok := pq.itemsm[key]
	if ok {
		return item.value
	}
	return nil
}

// readonly
func (pq *PrioQueue) GetIdx(idx int) PQItem {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	if len(pq.itemsv) > idx {
		itemin := pq.itemsv[idx]
		return itemin.value
	}
	return nil
}

// readonly
func (pq *PrioQueue) Last() PQItem {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	if len(pq.itemsv) > 0 {
		itemin := pq.itemsv[len(pq.itemsv)-1]
		return itemin.value
	}
	return nil
}

// readonly
func (pq *PrioQueue) Head(n ...int) []PQItem {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	items := []PQItem{}
	cnt := 1
	if len(n) > 0 {
		cnt = n[0]
	}
	if cnt <= 0 {
		return nil
	}

	for i := 0; i < cnt && i < len(pq.itemsv); i++ {
		itemin := pq.itemsv[i]
		items = append(items, itemin.value)
	}
	return items
}

// readonly
func (pq *PrioQueue) Tail(n ...int) []PQItem {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	items := []PQItem{}
	cnt := 1
	if len(n) > 0 {
		cnt = n[0]
	}
	if cnt <= 0 {
		return nil
	}

	for i := 0; i < cnt && i < len(pq.itemsv); i++ {
		idx := len(pq.itemsv) - i - 1
		itemin := pq.itemsv[idx]
		items = append(items, itemin.value)
	}
	return items
}

// readonly
func (pq *PrioQueue) Rand1() PQItem {
	items := pq.Randn(1)
	if len(items) > 0 {
		return items[0]
	}
	return nil
}
func (pq *PrioQueue) Randn(n ...int) []PQItem {
	pq.mu.RLock()
	defer pq.mu.RUnlock()
	items := []PQItem{}
	cnt := 1
	if len(n) > 0 {
		cnt = n[0]
	}
	if cnt <= 0 {
		return nil
	}

	if len(pq.itemsv) == 0 {
		return nil
	}
	rdidxes := RandNumsNodup(0, len(pq.itemsv), cnt)
	for _, idx := range rdidxes {
		itemin := pq.itemsv[idx]
		items = append(items, itemin.value)
	}

	return items
}
func (pq *PrioQueue) Range(iterfn func(item PQItem)) {
	pq.mu.RLock()
	items := make([]PQItem, len(pq.itemsv))
	for idx, item := range pq.itemsv {
		items[idx] = item.value
	}
	pq.mu.RUnlock()
	for _, item := range items {
		iterfn(item)
	}
}

///

// A PriorityQueue implements heap.Interface and holds Items.
type pqitemin struct {
	value PQItem
	index int
}
type priovec []*pqitemin

func (pq priovec) Len() int { return len(pq) }

func (pq priovec) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].value.Weight() > pq[j].value.Weight()
}

func (pq priovec) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priovec) Push(x interface{}) {
	n := len(*pq)
	item := x.(*pqitemin)
	item.index = n
	*pq = append(*pq, item)
	heap.Fix(pq, n)
}

func (pq *priovec) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
