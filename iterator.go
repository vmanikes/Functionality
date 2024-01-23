package functionality

// Iterator is the central backbone that used within this project
type Iterator[T any] []T

// NewIteratorWithLength Creates a new iterator with length of type T with 0 value types of type T
func NewIteratorWithLength[T any](len uint) Iterator[T] {
	return make(Iterator[T], len)
}

// NewIteratorWithCapacity Creates a new iterator with capacity of type T
func NewIteratorWithCapacity[T any](cap uint) Iterator[T] {
	return make(Iterator[T], 0, cap)
}

// NewIterator creates a new iterator without any len or cap
func NewIterator[T any]() Iterator[T] {
	return make(Iterator[T], 0)
}

func NewIteratorFromExistingSlice[T any](collection []T) Iterator[T] {
	iter := make(Iterator[T], 0, len(collection))

	for _, val := range collection {
		iter = append(iter, val)
	}

	return iter
}
