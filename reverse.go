package fp

// The Reverse() function creates a new slice populated
// with the values provided by the calling slice in reverse order
func Reverse[T any](arr []T) []T {
	lastIndex := len(arr) - 1
	result := make([]T, len(arr))
	reduce := func(acum []T, actual T, index int) []T {
		acum[lastIndex-index] = actual
		return acum
	}
	return Reduce(arr, reduce, result)
}
