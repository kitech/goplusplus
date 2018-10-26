package gods

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/maps/hashmap"
	// "github.com/emirpasic/gods/maps/treemap"
	"github.com/emirpasic/gods/sets/hashset"
	// "github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/emirpasic/gods/stacks/linkedliststack"
	// "github.com/emirpasic/gods/lists/doublylinkedlist"
)

func NewHashMap() *hashmap.Map { return hashmap.New() }

// func NewTreeMap() *treemap.Map { return treemap.New() }

func NewArrayList() *arraylist.List { return arraylist.New() }

func NewSinglyLinkedList() *singlylinkedlist.List { return singlylinkedlist.New() }

// func NewDoublyLinkedList() *doublylinkedlist.List { return do}

func NewHashSet() *hashset.Set { return hashset.New() }

// func NewTreeSet() *treeset.Set { return treeset.New() }

func NewLinkedListStack() *linkedliststack.Stack { return linkedliststack.New() }

func NewArrayStack() *arraystack.Stack { return arraystack.New() }
