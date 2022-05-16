package fp

// TODO: reuse ReduceI but first make some benchs
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

func ReduceI[T any, K any](arr []T, reducer func(K, T, int) K, initial K) (accumulator K) {
	if len(arr) == 0 {
		return initial
	}
	accumulator = initial
	for index, elem := range arr {
		accumulator = reducer(accumulator, elem, index)
	}
	return accumulator
}
