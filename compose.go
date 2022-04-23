package fp

func Compose[T any](funcs ...func(T) T) func(T) T {
	toCompose := Reverse(funcs)
	return Pipe(toCompose...)
}
