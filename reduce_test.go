package functionality

import (
	"github.com/matryer/is"
	"testing"
)

func TestIterator_Reduce_SumOfInts(t *testing.T) {
	length := 10

	iter := NewIteratorWithCapacity[int](uint(length))

	for i := 1; i <= length; i++ {
		iter = append(iter, i)
	}

	reducerFunc := func(accumulator, element int) int {
		return accumulator + element
	}

	result := iter.Reduce(reducerFunc, 0)

	is.New(t).Equal(result, 55)
}
