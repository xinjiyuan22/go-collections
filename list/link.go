package list

import (
	"sync"

	"github.com/xinjiyuan22/collections/collections"
)

type linkedNode[T collections.Object] struct {
	prev, next *linkedNode[T]
	value      *T
}

type LinkedList[T collections.Object] struct {
	sync.RWMutex
	head    *linkedNode[T]
	initCap int
}

func NewLinkedList[T collections.Object](cap int) collections.List[T] {
	return &LinkedList[T]{
		RWMutex: sync.RWMutex{},
		initCap: cap,
	}
}

func (a *LinkedList[T]) Add(o T) bool {
	a.Lock()
	defer a.Unlock()

	return true
}

func (a *LinkedList[T]) AddAll(c collections.Collection[T]) bool {
	return true
}

func (a *LinkedList[T]) Clear() {
	a.head = nil
}

func (a *LinkedList[T]) Contains(o T) bool {
	return true
}

func (a *LinkedList[T]) ContainsAll(c collections.Collection[T]) bool {
	return true
}

func (a *LinkedList[T]) IsEmpty() bool {
	a.RLock()
	defer a.RUnlock()
	return true
}

func (a *LinkedList[T]) Iterator() collections.Iterator[T] {
	return nil
}

func (a *LinkedList[T]) Remove(o T) bool {
	return true
}

func (a *LinkedList[T]) RemoveAll(c collections.Collection[T]) bool {
	return true
}

func (a *LinkedList[T]) ToArray() []T {
	a.RLock()
	defer a.RUnlock()
	res := make([]T, a.Size())
	return res
}

func (a *LinkedList[T]) Size() int {
	a.RLock()
	defer a.RUnlock()
	return 0
}

func (a *LinkedList[T]) Get(index int) *T {
	a.RLock()
	defer a.RUnlock()
	return nil
}

func (a *LinkedList[T]) Insert(index int, o T) {

}

func (a *LinkedList[T]) Set(index int, o T) {

}

func (a *LinkedList[T]) SubList(from, to int) collections.List[T] {
	return nil
}
