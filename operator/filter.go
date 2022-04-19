package operator

import (
	"github.com/xinjiyuan22/collections/collections"
)

// type FilterHandleFunc func[T collections.Object[any]](o T) bool

func Tes[T collections.Collection[collections.Object]]() {

}
func (c collections.Collection[T]) Filter(f FilterHandleFunc) collections.Collection[T] {

}
