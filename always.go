package fp

func Always[T any](elem T) func() T {
	return func() T {
		return elem
	}
}
