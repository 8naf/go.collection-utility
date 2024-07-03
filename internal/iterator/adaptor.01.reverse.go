package iterator

// Reverse returns an iterator that
// reverse Next and Prev behavior of the base iterator.
func Reverse[Index comparable, Value any](
	base BidirectionalIterator[Index, Value],
) BidirectionalIterator[Index, Value] {
	if r, ok := base.(*reversingIterator[Index, Value]); ok {
		return r.BidirectionalIterator
	}
	return &reversingIterator[Index, Value]{BidirectionalIterator: base}
}

type reversingIterator[Index comparable, Value any] struct {
	BidirectionalIterator[Index, Value]
}

func (iter *reversingIterator[Index, Value]) Next() bool {
	return iter.BidirectionalIterator.Prev()
}

func (iter *reversingIterator[Index, Value]) Prev() bool {
	return iter.BidirectionalIterator.Next()
}
