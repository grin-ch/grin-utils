package tool

type CustomKey[V any, T comparable] func(V) T

// MapFlip 翻转K-V的对应关系
// fn方法得到的key重复时只保留一个
func MapFlip[K, T comparable, V any](kvs map[K]V, fn CustomKey[V, T]) map[T]K {
	m := make(map[T]K, len(kvs))
	for k, v := range kvs {
		m[fn(v)] = k
	}
	return m
}

// MapFlipSlice 翻转K-V的对应关系
func MapFlipSlice[K, T comparable, V any](kvs map[K]V, fn CustomKey[V, T]) map[T][]K {
	m := make(map[T][]K, len(kvs))
	for k, v := range kvs {
		vals := m[fn(v)]
		vals = append(vals, k)
		m[fn(v)] = vals
	}
	return m
}

// MapKeys
func MapKeys[K comparable, V any](kvs map[K]V) []K {
	keys := make([]K, 0, len(kvs))
	for k := range kvs {
		keys = append(keys, k)
	}
	return keys
}

// MapValues
func MapValues[K comparable, V any](kvs map[K]V) []V {
	vals := make([]V, 0, len(kvs))
	for _, v := range kvs {
		vals = append(vals, v)
	}
	return vals
}
