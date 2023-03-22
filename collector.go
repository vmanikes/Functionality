package functionality

// Collect method is used to return the slice of type T is the users do not want an iterator type
func (iter Iterator[T]) Collect() []T {
	result := make([]T, 0, len(iter))

	for _, val := range iter {
		result = append(result, val)
	}

	return result
}
