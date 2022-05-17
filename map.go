package fp

// The Map() function creates a new slice populated with the results of calling
// a provided function on every element in the calling slice.
func Map[T any, K any](arr []T, mapper func(T) K) []K {
	reducer := func(acum []K, actual T, index int) []K {
		acum[index] = mapper(actual)
		return acum
	}
	initial := make([]K, len(arr))
	return Reduce(arr, reducer, initial)
}
