package stream

import (
	"github.com/xinjiyuan22/collections/collections"
	"github.com/xinjiyuan22/collections/list"
)

// type FilterHandleFunc func[T collections.Object[any]](o T) bool

func Filter[T collections.Object](c collections.Collection[T], f func(o *T) bool) collections.Collection[T] {
	it := c.Iterator()
	res := list.NewLinkedList[T]()
	for it.HasNext() {
		v := it.Next()
		if v == nil {
			panic("Got Nil Iterators")
		}
		if f(v) {
			res.Add(*v)
		}
	}
	return res
}
