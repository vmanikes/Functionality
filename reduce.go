// Package functionality contains the methods and objects to perform operations on go slices
package functionality

// ReducerFunc is a custom func that will be applied to items on an iterator
type ReducerFunc[T any] func(accumulator, element T) T

// Reduce is a receiver method on the Iterator that takes in a ReducerFunc and applies on the slice elements and returns
// a reduced value of type T
func (iter Iterator[T]) Reduce(reducerFunc ReducerFunc[T], initialValue T) T {
	accumulator := initialValue

	for _, val := range iter {
		accumulator = reducerFunc(accumulator, val)
	}

	return accumulator
}
