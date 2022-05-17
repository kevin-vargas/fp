package fp

// TODO: remove sideEffect from reduce function

// The Reverse() function creates a new slice populated
// with the values provided by the calling slice in reverse order
func Reverse[T any](arr []T) []T {
	index := 0
	lastIndex := len(arr) - 1
	result := make([]T, len(arr))
	reduce := func(acum []T, actual T) []T {
		acum[lastIndex-index] = actual
		index++
		return acum
	}
	return ReduceS(arr, reduce, result)
}
