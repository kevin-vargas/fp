package fp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reduce(t *testing.T) {
	// Arrange
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reducer := func(acum int, actual int) int {
		return acum + actual
	}

	// Act
	result := Reduce(arr, reducer, 0)

	// Assert
	expected := 45
	assert.Equal(t, expected, result)
}

func Test_Map(t *testing.T) {
	// Arrange
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mapper := func(num int) int {
		return num * 2
	}

	// Act
	result := Map(arr, mapper)

	// Assert
	expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	assert.Equal(t, expected, result)
}

func Test_Filter(t *testing.T) {
	// Arrange
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	predicate := func(num int) bool {
		return num > 5
	}

	// Act
	result := Filter(arr, predicate)

	// Assert
	expected := []int{6, 7, 8, 9}
	assert.Equal(t, expected, result)
}

func Test_Pipe(t *testing.T) {
	// Arrange
	toPipe := []func(int) int{
		func(i int) int {
			return i * 2
		},
		func(i int) int {
			return i + 3
		},
		func(i int) int {
			return i / 5
		},
	}
	value := 1

	// Act
	result := Pipe(toPipe...)(value)

	// Assert
	expected := 1
	assert.Equal(t, expected, result)
}

func Test_Reverse(t *testing.T) {
	// Arrange
	toReverse := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Act
	result := Reverse(toReverse)

	// Assert
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	assert.Equal(t, expected, result)
}

func Test_Compose(t *testing.T) {

	// Arrange
	toCompose := []func(int) int{
		func(i int) int {
			return i * 2
		},
		func(i int) int {
			return i + 3
		},
		func(i int) int {
			return i / 5
		},
	}
	value := 1

	// Act
	result := Compose(toCompose...)(value)

	// Assert
	expected := 6
	assert.Equal(t, expected, result)
}
