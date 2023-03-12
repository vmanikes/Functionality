package functionality

func (iter Iterator[T]) Map(mapFunc func(T) T) Iterator[T] {
	for idx, val := range iter {
		iter[idx] = mapFunc(val)
	}

	return iter
}
