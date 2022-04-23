package fp

func Pipe[T any](funcs ...func(T) T) func(T) T {
	return func(initial T) T {
		reduce := func(acum T, actual func(T) T) T {
			return actual(acum)
		}
		return Reduce(funcs, reduce, initial)
	}
}
