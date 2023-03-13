package functionality

import (
	"github.com/matryer/is"
	"testing"
)

func TestIterator_Collect(t *testing.T) {
	iter := NewIteratorWithCapacity[int](5)
	expectedResult := make([]int, 0, 5)

	for i := 1; i <= 5; i++ {
		iter = append(iter, i)
		expectedResult = append(expectedResult, i)
	}

	is.New(t).Equal(iter.Collect(), expectedResult)
}
