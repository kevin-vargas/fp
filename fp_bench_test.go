package fp

import (
	"strconv"
	"testing"
)

func Benchmark_Curry(b *testing.B) {
	fn := func(a int, b int) int {
		return a + b
	}
	fnCurried := Curry2(fn)
	b.Run("normal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fn(1, 3)
		}
	})
	b.Run("curried", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fnCurried(1)(3)
		}
	})

}

func Benchmark_Map(b *testing.B) {
	// Arrange
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mapper := func(num int) int {
		return num * 2
	}

	// Act
	b.Run("normal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := make([]int, len(arr))
			for i, elem := range arr {
				result[i] = mapper(elem)
			}
		}
	})

	// Act
	b.Run("mapped", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Map(arr, mapper)
		}
	})

}

func Benchmark_Reducer(b *testing.B) {
	// Arrange
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reducer := func(acum string, actual int) string {
		return acum + strconv.Itoa(actual)
	}
	reducerI := func(acum string, actual int, index int) string {
		return acum + strconv.Itoa(actual)
	}
	// Act
	b.Run("normal", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := ""
			for _, elem := range arr {
				str = str + strconv.Itoa(elem)
			}
		}
	})

	// Act
	b.Run("reducer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Reduce(arr, reducerI, "")
		}
	})
	// Act
	b.Run("reducerI", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ReduceS(arr, reducer, "")
		}
	})

}
