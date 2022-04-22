package stream

import (
	"github.com/xinjiyuan22/collections/collections"
	"github.com/xinjiyuan22/collections/list"
	"github.com/xinjiyuan22/collections/maps"
)

func MapToCollections[S collections.Object, T collections.Object](c collections.Collection[S], m func(o *S) T) collections.Collection[T] {
	res := list.NewArrayList[T](c.Size())
	iter := c.Iterator()
	for iter.HasNext() {
		v := iter.Next()
		if v == nil {
			panic("Invalid Containers")
		}

		res.Add(m(v))
	}
	return res
}

func MapToMaps[S collections.Object, K collections.Hashable, V collections.Object](c collections.Collection[S], m func(o *S) (K, V)) collections.Map[K, V] {
	res := maps.NewPlainMap[K, V]()
	iter := c.Iterator()

	for iter.HasNext() {
		v := iter.Next()
		if v == nil {
			panic("Invalid Containers")
		}

		res.Put(m(v))
	}
	return res
}
