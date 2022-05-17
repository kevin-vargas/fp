package fp

func ReduceS[T any, K any](arr []T, reducer func(K, T) K, initial K) (accumulator K) {
	r := func(k K, t T, i int) K {
		return reducer(k, t)
	}
	return Reduce(arr, r, initial)
}

func Reduce[T any, K any](arr []T, reducer func(K, T, int) K, initial K) (accumulator K) {
	if len(arr) == 0 {
		return initial
	}
	accumulator = initial
	for index, elem := range arr {
		accumulator = reducer(accumulator, elem, index)
	}
	return accumulator
}
