package iterator

import "github.com/8naf/go.collection-utility/common/typedef"

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
//
// This function is the variant of PrevN that supports all iterator types
// instead of only RandomAccessIterator.
func Prev[Index comparable, Value any](
	iter Iterator[Index, Value], n typedef.Quantity,
) bool {

	switch iter := (iter).(type) {
	case RandomAccessIterator[Index, Value]:
		return prevRAI(iter, n)
	case BidirectionalIterator[Index, Value]:
		return prevBI(iter, n)
	default:
		return prevI(iter, n)
	}
}

func prevRAI[Index comparable, Value any](
	iter RandomAccessIterator[Index, Value], n typedef.Quantity,
) bool {
	return iter.PrevN(n)
}

func prevBI[Index comparable, Value any](
	iter BidirectionalIterator[Index, Value], n typedef.Quantity,
) bool {
	return nextBI(iter, -n)
}

func prevI[Index comparable, Value any](
	iter Iterator[Index, Value], n typedef.Quantity,
) bool {
	return nextI(iter, -n)
}
