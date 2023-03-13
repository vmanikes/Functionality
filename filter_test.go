package functionality

import (
	"github.com/matryer/is"
	"testing"
)

func TestIterator_Filter_EvenNumbers(t *testing.T) {
	length := 10

	iter := NewIteratorWithCapacity[int](uint(length))
	expectedResult := []int{2, 4, 6, 8, 10}

	for i := 1; i <= length; i++ {
		iter = append(iter, i)
	}

	result := iter.Filter(func(i int) bool {
		return i%2 == 0
	}).Collect()

	is.New(t).Equal(result, expectedResult)
}
