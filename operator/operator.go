package operator

import (
	"github.com/xinjiyuan22/collections/collections"
)

type Operator[T collections.Object[any]] struct {
	collections.Collection[T]
}

func NewOperator[T collections.Object[any]](c collections.Collection[T]) *Operator[T] {
	return &Operator[T]{
		Collection: c,
	}
}
