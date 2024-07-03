package iterator

// Represents an iterator that can move through a collection in both directions.
type BidirectionalIterator[Index comparable, Value any] interface {
	Iterator[Index, Value]

	// Prev moves the iterator backward once if possible.
	//
	// If moving successfully,
	// returns true and
	// the iterator will move to the new position.
	// Otherwise,
	// returns false and
	// the iterator will retain the original position.
	Prev() bool
}
