package functionality

// MapFunc is a custom func that will be applied to items on an iterator
type MapFunc[T any] func(T) T

// Map is a receiver method on the Iterator that takes in a MapFunc and applies on the slice elements and returns a new Iterator
// with the filtered elements
func (iter Iterator[T]) Map(mapFunc MapFunc[T]) Iterator[T] {
	mapResult := NewIteratorWithCapacity[T](uint(len(iter)))

	for _, val := range iter {
		mapResult = append(mapResult, mapFunc(val))
	}

	return mapResult
}
