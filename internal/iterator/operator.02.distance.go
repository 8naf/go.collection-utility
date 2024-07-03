package iterator

import (
	"github.com/8naf/go.collection-utility/common/constant"
	"github.com/8naf/go.collection-utility/common/typedef"
)

// If `to` is `ToLowerBound`, Distance counts the remaining number of times
// `Prev` method can be called until it returns false.
//
// If `to` is `ToUpperBound`, Distance counts the remaining number of times
// `Next` method can be called until it returns false.
//
// Returns a negative number
// if that number of times is infinite,
// or if any errors occur.
//
// This function is the variant of Distance that supports all iterator types
// instead of only RandomAccessIterator.
func Distance[Index comparable, Value any](
	from Iterator[Index, Value], to Direction,
) (count typedef.Quantity) {
	switch from := from.(type) {
	case RandomAccessIterator[Index, Value]:
		return distanceRAI(from, to)
	case BidirectionalIterator[Index, Value]:
		return distanceBI(from, to)
	default:
		return distanceI(from, to)
	}
}

func distanceRAI[Index comparable, Value any](
	from RandomAccessIterator[Index, Value], to Direction,
) (count typedef.Quantity) {
	return from.Distance(to)
}

func distanceBI[Index comparable, Value any](
	from BidirectionalIterator[Index, Value], to Direction,
) (count typedef.Quantity) {
	from = Clone[Index, Value](from)

	move := func() bool {
		return from.Next()
	}
	if to == ToLowerBound {
		move = func() bool {
			return from.Prev()
		}
	}

	for count = 0; count >= 0 && move(); count++ {
	}

	if count < 0 {
		count = constant.INFINITY
	}
	return
}

func distanceI[Index comparable, Value any](
	from Iterator[Index, Value], to Direction,
) (count typedef.Quantity) {
	if to == ToLowerBound {
		return constant.OUT_OF_RANGE
	}

	from = Clone[Index, Value](from)
	for count = 0; count >= 0 && from.Next(); count++ {
	}

	if count < 0 {
		count = constant.INFINITY
	}
	return
}

// #endregion Distance
