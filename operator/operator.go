package operator

import (
	"github.com/xinjiyuan22/collections/collections"
)

type Operator[T collections.Object] struct {
	collections.Collection[T]
}

func NewOperator[T collections.Object](c collections.Collection[T]) *Operator[T] {
	return &Operator[T]{
		Collection: c,
	}
}
