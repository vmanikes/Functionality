package functionality

import (
	"fmt"
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

func Boo() {
	numbers := NewIterator[int]()

	for i := 1; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	sum := numbers.Filter(func(num int) bool {
		return num%2 == 0
	}).Map(func(i int) int {
		return i * i
	}).Reduce(func(accumulator, element int) int {
		return accumulator + element
	}, 0)

	fmt.Println("SUM IS ", sum)
}