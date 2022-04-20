package operator

func (c *Operator[T]) Sort(less func(x, y T) bool) *Operator[T] {
	return c
}
