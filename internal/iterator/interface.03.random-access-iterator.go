package iterator

import "github.com/8naf/go.collection-utility/common/typedef"

// Represents a bidirectional iterator that can move through a collection
// with an arbitrary offset instead of one at a time.
type RandomAccessIterator[Index comparable, Value any] interface {
	BidirectionalIterator[Index, Value]

	// If possible, NextN moves the iterator n times
	// forward if n is positive,
	// or backward if n is negative.
	//
	// If moving successfully,
	// returns true and
	// the iterator will move to the new position.
	// Otherwise,
	// returns false and
	// the iterator will retain the original position.
	NextN(n typedef.Quantity) bool

	// If possible, PrevN moves the iterator n times
	// backward if n is positive,
	// or forward if n is negative.
	//
	// If moving successfully,
	// returns true and
	// the iterator will move to the new position.
	// Otherwise,
	// returns false and
	// the iterator will retain the original position.
	PrevN(n typedef.Quantity) bool

	// If `to` is `ToLowerBound`, Distance counts the remaining number of times
	// `Prev` method can be called until it returns false.
	//
	// If `to` is `ToUpperBound`, Distance counts the remaining number of times
	// `Next` method can be called until it returns false.
	//
	// Returns a negative number
	// if that number of times is infinite,
	// or if any errors occur.
	Distance(to Direction) typedef.Quantity
}

type Direction int

const (
	ToLowerBound Direction = -1
	ToUpperBound Direction = 1
)
