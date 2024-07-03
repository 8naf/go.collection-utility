package iterator

// WithExitCondition returns an iterator that stops
// when `exitCondition` returns true.
func WithExitCondition[Index comparable, Value any](
	iterator Iterator[Index, Value],
	exitCondition func(Iterator[Index, Value]) bool,
) Iterator[Index, Value] {
	return &withExitConditionIterator[Index, Value]{iterator, exitCondition}
}

type withExitConditionIterator[Index comparable, Value any] struct {
	Iterator[Index, Value]
	exitCondition func(Iterator[Index, Value]) bool
}

func (iter *withExitConditionIterator[Index, Value]) Next() bool {
	return iter.Iterator.Next() && !iter.exitCondition(iter.Iterator)
}
