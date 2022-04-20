package operator

func (c *Operator[T]) ForEach(f func(index int, o T)) *Operator[T] {
	it := c.Iterator()

	for it.HasNext() {

	}
	return c
}
