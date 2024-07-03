package math

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

func Sign[T Signed](num T) T {
	if num == 0 {
		return 0
	}
	if num > 0 {
		return 1
	}
	return -1
}

func Abs[T Integer](num T) T {
	if num < 0 {
		return -num
	}
	return num
}
