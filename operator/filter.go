package operator

// type FilterHandleFunc func[T collections.Object[any]](o T) bool

func (c Operator[T]) Filter(f func(o T) bool) Operator[T] {
	it := c.Iterator()

	for it.HasNext() {

	}
	return c
}
