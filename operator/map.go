package operator

import (
	"github.com/golang/protobuf/ptypes/any"
	"github.com/xinjiyuan22/collections/collections"
)

func (c Operator[T]) MapToOperator(f func(o T) collections.Map[S collections.Hashable, V collections.Object[any]]) {
	it := c.Iterator()

	for it.HasNext() {

	}

}
