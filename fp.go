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

func Map[T any, K any](arr []T, mapper func(T) K) []K {
	reducer := func(acum []K, actual T) []K {
		acum = append(acum, mapper(actual))
		return acum
	}
	initial := make([]K, 0)
	return Reduce(arr, reducer, initial)
}

func Filter[T any](arr []T, predicate func(T) bool) []T {
	reducer := func(acum []T, actual T) []T {
		if predicate(actual) {
			acum = append(acum, actual)
		}
		return acum
	}
	initial := make([]T, 0)
	return Reduce(arr, reducer, initial)
}
