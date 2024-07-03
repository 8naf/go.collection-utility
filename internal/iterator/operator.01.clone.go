package iterator

// The clone method of an iterator sometimes needs type assertion
// to accurately returns the type of the original iterator.
// This function serves to perform that step.
//
// It will panic if the type of the clone
// is not the same as the original iterator.
func Clone[
	Index comparable,
	Value any,
	Iter Iterator[Index, Value],
](iter Iter) Iter {
	clone, ok := iter.Clone().(Iter)
	if !ok {
		panic("type of clone must be the same as the original iterator")
	}

	return clone
}
