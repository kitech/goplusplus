package gopp

import (
	"container/heap"
	"math"
	"sync"
)

// weight小，优先级高
type PQItem interface {
	Key() interface{}
	Weight() int
	// Less(that PQItem) bool
}

// mixed queue and map features
type PrioQueue struct {
	mu     sync.RWMutex
	itemsv priovec
	itemsm map[interface{}]*pqitemin
	maxcnt int
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

	olditem, ok := pq.itemsm[item.Key()]
	if ok {
		_ = olditem
		// olditem.value = item
		// heap.Fix(&pq.itemsv, olditem.index)
	} else {
		itemin := &pqitemin{value: item}
		heap.Push(&pq.itemsv, itemin)
		pq.itemsm[item.Key()] = itemin
	}
}
func (pq *PrioQueue) PushSet(item PQItem) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	olditem, ok := pq.itemsm[item.Key()]
	if ok {
		_ = olditem
		olditem.value = item
		heap.Fix(&pq.itemsv, olditem.index)
	} else {
		itemin := &pqitemin{value: item}
		heap.Push(&pq.itemsv, itemin)
		pq.itemsm[item.Key()] = itemin
	}
}

func (pq *PrioQueue) Pop() PQItem {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	itemx := heap.Pop(&pq.itemsv)
	item := itemx.(*pqitemin)
	delete(pq.itemsm, item.value.Key())
	return item.value
}

func (pq *PrioQueue) Len() int {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	Assert(len(pq.itemsv) == len(pq.itemsm), "Ooops, invalid inner state", len(pq.itemsv), len(pq.itemsm))
	return len(pq.itemsv)
}

func (pq *PrioQueue) Del(key interface{}) PQItem {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	item, ok := pq.itemsm[key]
	if ok {
		heap.Remove(&pq.itemsv, item.index)
		return item.value
	}
	return nil
}

// readonly
func (pq *PrioQueue) GetKey(key interface{}) PQItem {
	pq.mu.RLock()
	defer pq.mu.RUnlock()

	item, _ := pq.itemsm[key]
	return item.value
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
	for i := 0; i < cnt && i < len(pq.itemsv); i++ {
		idx := len(pq.itemsv) - i - 1
		itemin := pq.itemsv[idx]
		items = append(items, itemin.value)
	}
	return items
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
	return pq[i].value.Weight() < pq[j].value.Weight()
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
