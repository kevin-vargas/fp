package fp

import "math"

func Chunk[T any](size int, arr []T) [][]T {
	sizeResult := int(math.Ceil(float64(len(arr)) / float64(size)))
	remanent := len(arr) % size
	result := make([][]T, sizeResult)
	reducer := func(acum [][]T, actual T, index int) [][]T {
		indexX := index / size
		if acum[indexX] == nil {
			sizeOfSlice := size
			if indexX == sizeResult-1 && remanent != 0 {
				sizeOfSlice = remanent
			}
			acum[indexX] = make([]T, sizeOfSlice)
		}
		indexY := index % size
		acum[indexX][indexY] = actual
		return acum
	}
	return ReduceI(arr, reducer, result)
}
