package iterator

// Represents an object that can move through a collection in one direction.
type Iterator[Index comparable, Value any] interface {

	// Clone returns a clone of the iterator.
	//
	// The returned clone must:
	// refer to the same collection as the original iterator;
	// have the same position on the collection as the original iterator; and
	// have the same underlying type as the original iterator.
	//
	// Cloning the metadata and any other state of the original iterator
	// is optional.
	//
	// Moving the clone must not affect the original iterator's state.
	Clone() Iterator[Index, Value]

	// Index returns the index of the element
	// at the iterator's current position.
	//
	// Returns a negative number
	// if the iterator is at an invalid position.
	//
	// If the iterator has not moved yet,
	// then its position is considered invalid,
	// and Index must return `OUT_OF_RANGE`.
	Index() Index

	// Value returns the value of the element
	// at the iterator's current position.
	//
	// Returns the zero value of `Value` type
	// if the iterator is at an invalid position.
	//
	// If the iterator has not moved yet,
	// then its position is considered invalid,
	// and Value must return the zero value of `Value` type.
	Value() Value

	// Metadata returns the metadata of the iterator.
	Metadata() Metadata

	// Next moves the iterator forward once if possible.
	//
	// If moving successfully,
	// returns true and
	// the iterator will move to the new position.
	// Otherwise,
	// returns false and
	// the iterator will retain the original position.
	Next() bool
}

type Metadata = map[string]any
