package stream

import "github.com/xinjiyuan22/collections/collections"

func Min[T collections.Comparable](c collections.Collection[T]) *T {
	iter := c.Iterator()
	var res *T
	for iter.HasNext() {
		v := iter.Next()
		if v == nil {
			panic("Get Invalid Containers")
		}
		if res == nil {
			res = v
		}

		if (*res).Less(*v) {
			res = v
		}
	}
	return res
}
func Max[T collections.Comparable](c collections.Collection[T]) *T {
	iter := c.Iterator()
	var res *T
	for iter.HasNext() {
		v := iter.Next()
		if v == nil {
			panic("Get Invalid Containers")
		}
		if res == nil {
			res = v
		}

		if (*v).Less(*res) {
			res = v
		}
	}
	return res
}
