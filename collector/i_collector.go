package collector

func newEachOpt(opts ...EachOpt) *eachOpt {
	eo := new(eachOpt)
	for _, fn := range opts {
		fn(eo)
	}
	return eo
}

type eachOpt struct {
	reverse bool
}

type EachOpt func(*eachOpt)

func OptReverse() EachOpt { return func(eo *eachOpt) { eo.reverse = true } }

type Collector[T comparable] interface {
	Len() int
	IsEmpty() bool
	Slice(...EachOpt) []T
	Add(T)
	Remove() T
}

type Iterator[T comparable] interface {
	ForEach(func(T) bool, ...EachOpt)
}
