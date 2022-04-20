package list

import (
	"sync"

	"github.com/xinjiyuan22/collections/collections"
)

type LinkIterator[T collections.Object] struct {
	array   *LinkedList[T]
	current *linkedNode[T]
}

func (it *LinkIterator[T]) HasNext() bool {
	return it.current.next != nil
}

func (it *LinkIterator[T]) Next() *T {
	if !it.HasNext() {
		return nil
	}
	it.current = it.current.next
	return it.current.value
}

type linkedNode[T collections.Object] struct {
	prev, next *linkedNode[T]
	value      *T
}

type LinkedList[T collections.Object] struct {
	sync.RWMutex
	head *linkedNode[T]
	tail *linkedNode[T]
	size int
}

func NewLinkedList[T collections.Object]() collections.List[T] {
	res := &LinkedList[T]{
		RWMutex: sync.RWMutex{},
		head:    &linkedNode[T]{},
	}
	res.tail = res.head
	return res
}

func NewLinkedListFromSlice[T collections.Object](s []*T) collections.List[T] {
	res := NewLinkedList[T]()
	for _, v := range s {
		res.Add(*v)
	}
	return res
}

func (a *LinkedList[T]) Add(o T) bool {
	a.Lock()
	defer a.Unlock()
	a.tail.next = new(linkedNode[T])
	a.tail.next.prev = a.tail
	a.tail.next.value = &o
	a.tail = a.tail.next
	a.size++
	return true
}

func (a *LinkedList[T]) AddAll(c collections.Collection[T]) bool {
	iter := c.Iterator()
	for iter.HasNext() {
		a.Add(*iter.Next())
	}
	return true
}

func (a *LinkedList[T]) Clear() {
	a.head = &linkedNode[T]{}
	a.tail = a.head
	a.size = 0
}

func (a *LinkedList[T]) Contains(o T) bool {
	a.RLock()
	defer a.RUnlock()
	iter := a.Iterator()
	for iter.HasNext() {
		if (*iter.Next()).Equal(o) {
			return true
		}
	}
	return false
}

func (a *LinkedList[T]) ContainsAll(c collections.Collection[T]) bool {
	res := true
	iter := c.Iterator()

	for iter.HasNext() {
		v := a.Contains(*iter.Next())
		res = res && v
	}
	return res
}

func (a *LinkedList[T]) IsEmpty() bool {
	a.RLock()
	defer a.RUnlock()
	return a.head == a.tail
}

func (a *LinkedList[T]) Iterator() collections.Iterator[T] {
	return &LinkIterator[T]{
		array:   a,
		current: a.head,
	}
}

func (a *LinkedList[T]) Remove(o T) bool {
	var res bool
	if a.IsEmpty() {
		return false
	}
	a.Lock()
	defer a.Unlock()
	current := a.head.next
	for current.next != nil {
		if (*current.value).Equal(o) {
			current.prev.next = current.next
			current.next.prev = current.prev
			res = true
			a.size--
		}
		current = current.next
	}

	if (*current.value).Equal(o) {
		current.prev.next = nil
		a.size--
		res = true
	}
	return res
}

func (a *LinkedList[T]) RemoveAll(c collections.Collection[T]) bool {
	var res bool = false
	iter := c.Iterator()
	for iter.HasNext() {
		r := a.Remove(*iter.Next())
		res = res || r
	}
	return res
}

func (a *LinkedList[T]) ToArray() []T {
	a.RLock()
	defer a.RUnlock()
	res := make([]T, a.Size())
	iter := a.Iterator()
	var i int
	for iter.HasNext() {
		res[i] = *iter.Next()
		i++
	}
	return res
}

func (a *LinkedList[T]) Size() int {
	a.RLock()
	defer a.RUnlock()
	return a.size
}

func (a *LinkedList[T]) Get(index int) *T {
	a.RLock()
	defer a.RUnlock()
	if index >= a.size {
		return nil
	}

	iter := a.Iterator()
	for iter.HasNext() && index > 0 {
		iter.Next()
		index--
	}
	return iter.Next()
}

func (a *LinkedList[T]) Insert(index int, o T) {
	if index >= a.Size() {
		return
	}
	a.Lock()
	defer a.Unlock()
	current := a.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	node := &linkedNode[T]{
		value: &o,
	}

	node.next = current.next
	current.prev = node
	node.prev = current.prev
	current.next = node
}

func (a *LinkedList[T]) Set(index int, o T) {
	if index >= a.Size() {
		return
	}

	a.Lock()
	defer a.Unlock()
	current := a.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	current.value = &o
}

func (a *LinkedList[T]) SubList(from, to int) collections.List[T] {
	if from > to || to >= a.Size() {
		return nil
	}
	a.RLock()
	defer a.RUnlock()
	res := NewLinkedList[T]()
	current := a.head
	for i := 0; i < from; i++ {
		current = current.next
	}

	for i := from; i < to; i++ {
		res.Add(*current.value)
		current = current.next
	}
	return res
}
