package set

import (
	"github.com/xinjiyuan22/collections/collections"
)

type OrderSet[T collections.Comparable] struct {
}

func NewOrderSet[T collections.Comparable]() collections.Set[T] {
	return &OrderSet[T]{}
}

func NewOrderSetWithContainer[T collections.Comparable](c collections.Collection[T]) collections.Set[T] {
	return &OrderSet[T]{}
}

func (s *OrderSet[T]) Add(o T) bool {
	return false
}

func (s *OrderSet[T]) AddAll(c collections.Collection[T]) bool {
	return false
}

func (s *OrderSet[T]) Clear() {
}

func (s *OrderSet[T]) Contains(o T) bool {
	return false
}

func (s *OrderSet[T]) ContainsAll(c collections.Collection[T]) bool {
	return false
}

func (s *OrderSet[T]) IsEmpty() bool {
	return false
}

func (s *OrderSet[T]) Iterator() collections.Iterator[T] {
	return nil
}

func (s *OrderSet[T]) Remove(o T) bool {
	return false
}

func (s *OrderSet[T]) RemoveAll(c collections.Collection[T]) bool {
	return false
}

func (s *OrderSet[T]) ToArray() []T {
	return nil
}

func (s *OrderSet[T]) Size() int {
	return 0
}
