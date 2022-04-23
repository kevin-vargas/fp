package fp

func Reduce[T any, K any](arr []T, reducer func(K, T) K, initial K) (accumulator K) {
	if len(arr) == 0 {
		return initial
	}
	accumulator = initial
	for _, elem := range arr {
		accumulator = reducer(accumulator, elem)
	}
	return accumulator
}
