package list

import (
	"sync"

	"github.com/xinjiyuan22/collections/collections"
)

type ArrayIterator[T collections.Object] struct {
	array        *ArrayList[T]
	currentIndex int
}

func (it *ArrayIterator[T]) HasNext() bool {
	return it.currentIndex < it.array.Size()-1
}

func (it *ArrayIterator[T]) Next() *T {
	if !it.HasNext() {
		return nil
	}
	it.currentIndex++
	return it.array.Get(it.currentIndex)
}

type ArrayList[T collections.Object] struct {
	sync.RWMutex
	content []*T
	initCap int
}

func NewArrayList[T collections.Object](cap int) collections.List[T] {
	return &ArrayList[T]{
		RWMutex: sync.RWMutex{},
		content: make([]*T, 0, cap),
		initCap: cap,
	}
}

func NewArrayListFromSlice[T collections.Object](s []*T) collections.List[T] {
	return &ArrayList[T]{
		RWMutex: sync.RWMutex{},
		content: s,
		initCap: len(s),
	}
}

func (a *ArrayList[T]) Add(o T) bool {
	a.Lock()
	defer a.Unlock()
	a.content = append(a.content, &o)
	return true
}

func (a *ArrayList[T]) AddAll(c collections.Collection[T]) bool {
	a.Lock()
	defer a.Unlock()
	it := c.Iterator()
	for it.HasNext() {
		a.content = append(a.content, it.Next())
	}
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

func (a *ArrayList[T]) ContainsAll(c collections.Collection[T]) bool {
	it := c.Iterator()
	for it.HasNext() {
		if !a.Contains(*it.Next()) {
			return false
		}
	}
	return true
}

func (a *ArrayList[T]) IsEmpty() bool {
	a.RLock()
	defer a.RUnlock()
	return len(a.content) == 0
}

func (a *ArrayList[T]) Iterator() collections.Iterator[T] {
	return &ArrayIterator[T]{
		array:        a,
		currentIndex: -1,
	}
}

func (a *ArrayList[T]) Remove(o T) bool {
	a.Lock()
	defer a.Unlock()
	var i, j int
	for i, j = 0, 0; i < len(a.content); i++ {
		if (*a.content[i]).Equal(o) {
			continue
		}
		a.content[j] = a.content[i]
		j++
	}

	if j == i {
		return false
	}
	a.content = a.content[:j]
	return true
}

func (a *ArrayList[T]) RemoveAll(c collections.Collection[T]) bool {
	iter := c.Iterator()
	var res bool
	for iter.HasNext() {
		c := a.Remove(*iter.Next())
		res = res || c
	}
	return res
}

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
	if len(a.content) <= index {
		return nil
	}
	return a.content[index]
}

func (a *ArrayList[T]) Insert(index int, o T) {
	a.Lock()
	defer a.Unlock()
	size := len(a.content)
	if index > len(a.content) || index < 0 {
		return
	}
	a.content = append(a.content, &o)
	tmp := a.content[size]
	for i := size; i > index; i-- {
		a.content[i] = a.content[i-1]
	}
	a.content[index] = tmp
}

func (a *ArrayList[T]) Set(index int, o T) {
	if index >= a.Size() {
		return
	}
	a.Lock()
	defer a.Unlock()
	a.content[index] = &o

}

func (a *ArrayList[T]) SubList(from, to int) collections.List[T] {
	if from > to || to < a.Size() {
		return nil
	}

	res := NewArrayList[T](to - from)
	for i := from; i < to; i++ {
		res.Add(*a.Get(i))
	}
	return res
}
