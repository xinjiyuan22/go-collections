package list

import (
	"sync"

	"github.com/xinjiyuan22/collections/collections"
)

type linkedNode[T collections.Object[any]] struct {
	prev, next *linkedNode[T]
	value      *T
}

type LinkedList[T collections.Object[any]] struct {
	sync.RWMutex
	head    *linkedNode[T]
	initCap int
}

func NewLinkedList[T collections.Object[any]](cap int) collections.List[T] {
	return &LinkedList[T]{
		RWMutex: sync.RWMutex{},
		content: make([]*T, 0, cap),
		initCap: cap,
	}
}

func (a *LinkedList[T]) Add(index int, o T) bool {
	a.Lock()
	defer a.Unlock()

	return true
}

func (a *LinkedList[T]) AddAll(c collections.Collection[T]) bool {
	return true
}

func (a *LinkedList[T]) Clear() {
	a.Lock()
	defer a.Unlock()
	a.content = make([]*T, 0, a.initCap)
}

func (a *LinkedList[T]) Contains(o T) bool {
	a.RLock()
	defer a.RUnlock()
	for _, t := range a.content {
		if (*t).Equal(o) {
			return true
		}
	}
	return false
}

func (a *LinkedList[T]) ContainsAll(c collections.Collection[T]) bool

func (a *LinkedList[T]) IsEmpty() bool {
	a.RLock()
	defer a.RUnlock()
	return len(a.content) == 0
}

func (a *LinkedList[T]) Iterator() *collections.Iterator[T]

func (a *LinkedList[T]) Remove(o T) bool

func (a *LinkedList[T]) RemoveAll(c collections.Collection[T]) bool

func (a *LinkedList[T]) ToArray() []T {
	a.RLock()
	defer a.RUnlock()
	res := make([]T, a.Size())
	for i, t := range a.content {
		res[i] = *t
	}
	return res
}

func (a *LinkedList[T]) Size() int {
	a.RLock()
	defer a.RUnlock()
	return len(a.content)
}

func (a *LinkedList[T]) Get(index int) *T {
	a.RLock()
	defer a.RUnlock()
	if len(a.content) >= index {
		return nil
	}
	return a.content[index]
}

func (a *LinkedList[T]) Append(o ...T) {

}

func (a *LinkedList[T]) Set(index int, o T) {

}

func (a *LinkedList[T]) SubList(from, to int) collections.List[T] {
	return nil
}
