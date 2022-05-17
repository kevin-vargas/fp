package fp

// The filter() function creates a new slice with all elements
// that pass the test implemented by the provided function.
func Filter[T any](arr []T, predicate func(T) bool) []T {
	reducer := func(acum []T, actual T) []T {
		if predicate(actual) {
			acum = append(acum, actual)
		}
		return acum
	}
	initial := make([]T, 0)
	return ReduceS(arr, reducer, initial)
}
