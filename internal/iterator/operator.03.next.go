package iterator

import (
	"reflect"

	"github.com/8naf/go.collection-utility/common/typedef"
)

// If possible, Next moves the iterator n times
// forward if n is positive,
// or backward if n is negative.
//
// If moving successfully,
// returns true and
// the iterator will move to the new position.
// Otherwise,
// returns false and
// the iterator will retain the original position.
//
// This function is the variant of NextN that supports all iterator types
// instead of only RandomAccessIterator.
func Next[Index comparable, Value any](
	iter Iterator[Index, Value], n typedef.Quantity,
) bool {
	switch iter := (iter).(type) {
	case RandomAccessIterator[Index, Value]:
		return nextRAI(iter, n)
	case BidirectionalIterator[Index, Value]:
		return nextBI(iter, n)
	default:
		return nextI(iter, n)
	}
}

func nextRAI[Index comparable, Value any](
	iter RandomAccessIterator[Index, Value], n typedef.Quantity,
) bool {
	return iter.NextN(n)
}

func nextBI[Index comparable, Value any](
	iter BidirectionalIterator[Index, Value], n typedef.Quantity,
) bool {
	clone := Clone[Index, Value](iter)

	canMove := func() bool {
		if n != 0 {
			n--
		}
		return clone.Next()
	}
	if n < 0 {
		canMove = func() bool {
			if n != 0 {
				n++
			}
			return clone.Prev()
		}
	}

	for canMove() && n != 0 {
	}

	if n != 0 {
		return false
	}

	set[Index, Value](iter, clone)
	return true
}

func nextI[Index comparable, Value any](
	iter Iterator[Index, Value], n typedef.Quantity,
) bool {
	if n < 0 {
		return false
	}

	clone := Clone[Index, Value](iter)
	for ; n != 0 && clone.Next(); n-- {
	}

	if n != 0 {
		return false
	}

	set(iter, clone)
	return true
}

func set[Index comparable, Value any](
	iter Iterator[Index, Value], other Iterator[Index, Value],
) {
	vIter := reflect.ValueOf(iter)
	vOther := reflect.ValueOf(other)

	vIter.Elem().Set(vOther.Elem())
}
