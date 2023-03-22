package functionality

// FilterFunc is a custom func that will be applied to items on an iterator
type FilterFunc[T any] func(T) bool

// Filter is a receiver method on the Iterator that takes in a FilterFunc and applies on the slice elements and returns a new Iterator
// with the filtered elements
func (iter Iterator[T]) Filter(filterFunc FilterFunc[T]) Iterator[T] {
	filteredIterator := NewIterator[T]()

	for _, val := range iter {
		if filterFunc(val) {
			filteredIterator = append(filteredIterator, val)
		}
	}

	return filteredIterator
}
