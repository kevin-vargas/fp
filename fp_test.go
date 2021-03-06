package fp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReduceS(t *testing.T) {
	// Arrange
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reducer := func(acum int, actual int) int {
		return acum + actual
	}

	// Act
	result := ReduceS(arr, reducer, 0)

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

func Test_Identity(t *testing.T) {
	// Arrange
	cases := []struct {
		desc   string
		value  any
		expect any
	}{
		{
			desc:   "string",
			value:  "test",
			expect: "test",
		},
		{
			desc:   "int",
			value:  3,
			expect: 3,
		},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			// Act
			result := Identity(tt.value)

			// Assert
			assert.Equal(t, tt.expect, result)
		})
	}
}

func Test_Always(t *testing.T) {
	// Arrange
	cases := []struct {
		desc   string
		value  any
		expect any
	}{
		{
			desc:   "string",
			value:  "test",
			expect: "test",
		},
		{
			desc:   "int",
			value:  3,
			expect: 3,
		},
	}
	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			// Act
			result := Always(tt.value)

			// Assert
			assert.Equal(t, tt.expect, result())
		})
	}
}

func Test_Curry(t *testing.T) {
	// Arrange
	fn := func(a int, b int) int {
		return a + b
	}

	//Act
	fnCurry := Curry2(fn)
	result := fnCurry(1)(2)

	//Assert
	assert.Equal(t, 3, result)
}

func Test_Chunk(t *testing.T) {
	// Arrange
	arr := []int{1, 2, 3, 4, 5}
	cases := []struct {
		size   int
		expect [][]int
	}{
		{
			1,
			[][]int{{1}, {2}, {3}, {4}, {5}},
		},
		{
			2,
			[][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			3,
			[][]int{{1, 2, 3}, {4, 5}},
		},
		{
			4,
			[][]int{{1, 2, 3, 4}, {5}},
		},
		{
			5,
			[][]int{{1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprintf("with size of %v", tt.size), func(t *testing.T) {
			// Act
			result := Chunk(tt.size, arr)

			// Assert
			assert.ElementsMatch(t, tt.expect, result)
		})
	}
}
