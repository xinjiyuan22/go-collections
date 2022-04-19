package list

import (
	"sync"

	"github.com/xinjiyuan22/collections/collections"
)

type ArrayList[T collections.Object[any]] struct {
	sync.RWMutex
	content []*T
	initCap int
}

func NewArrayList[T collections.Object[any]](cap int) collections.List[T] {
	return &ArrayList[T]{
		RWMutex: sync.RWMutex{},
		content: make([]*T, 0, cap),
		initCap: cap,
	}
}

func (a *ArrayList[T]) Add(index int, o T) bool {
	a.Lock()
	defer a.Unlock()

	return true
}

func (a *ArrayList[T]) AddAll(c collections.Collection[T]) bool {
	return true
}

func (a *ArrayList[T]) Clear() {
	a.Lock()
	defer a.Unlock()
	a.content = make([]*T, 0, a.initCap)
}

func (a *ArrayList[T]) Contains(o T) bool {
	a.RLock()
	defer a.RUnlock()
	for _, t := range a.content {
		if (*t).Equal(o) {
			return true
		}
	}
	return false
}

func (a *ArrayList[T]) ContainsAll(c collections.Collection[T]) bool

func (a *ArrayList[T]) IsEmpty() bool {
	a.RLock()
	defer a.RUnlock()
	return len(a.content) == 0
}

func (a *ArrayList[T]) Iterator() *collections.Iterator[T]

func (a *ArrayList[T]) Remove(o T) bool

func (a *ArrayList[T]) RemoveAll(c collections.Collection[T]) bool

func (a *ArrayList[T]) ToArray() []T {
	a.RLock()
	defer a.RUnlock()
	res := make([]T, a.Size())
	for i, t := range a.content {
		res[i] = *t
	}
	return res
}

func (a *ArrayList[T]) Size() int {
	a.RLock()
	defer a.RUnlock()
	return len(a.content)
}

func (a *ArrayList[T]) Get(index int) *T {
	a.RLock()
	defer a.RUnlock()
	if len(a.content) >= index {
		return nil
	}
	return a.content[index]
}

func (a *ArrayList[T]) Append(o ...T) {

}

func (a *ArrayList[T]) Set(index int, o T) {

}

func (a *ArrayList[T]) SubList(from, to int) collections.List[T] {
	return nil
}
