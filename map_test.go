package functionality

import (
	"github.com/matryer/is"
	"testing"
)

func TestIterator_Map_SquareOfInput(t *testing.T) {
	length := 10

	iter := NewIteratorWithCapacity[int](uint(length))
	expectedResult := make([]int, 0, length)

	for i := 1; i <= length; i++ {
		iter = append(iter, i)
		expectedResult = append(expectedResult, i*i)
	}

	result := iter.Map(func(i int) int {
		return i * i
	}).Collect()

	is.New(t).Equal(result, expectedResult)
}
