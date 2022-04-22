package stream

import "github.com/xinjiyuan22/collections/collections"

func ForEach[T collections.Object](c collections.Collection[T], f func(index int, o *T)) {
	it := c.Iterator()
	index := 0
	for it.HasNext() {
		v := it.Next()
		if v == nil {
			panic("Get Invalid Iterator")
		}
		f(index, v)
		index++
	}
}
