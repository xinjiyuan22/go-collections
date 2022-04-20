package operator

func (c *Operator[T]) Distinct() *Operator[T] {
	it := c.Iterator()

	for it.HasNext() {

	}
	return c
}
