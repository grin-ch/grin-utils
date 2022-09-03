package math

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Uint | Int | Float
}

// Comparator is not comparable
type Comparator interface {
	Number | ~string
}

func Max[T Comparator](a, b T) T {
	if b > a {
		return b
	}
	return a
}

func Min[T Comparator](a, b T) T {
	if b < a {
		return b
	}
	return a
}

func Abs[T Number](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
