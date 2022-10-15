package collector

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type T int

func TestNewLinkList(t *testing.T) {
	list := NewLinkList[T]()
	assert.Empty(t, list)
}

func TestEmpty(t *testing.T) {
	list := NewLinkList[T]()
	assert.True(t, list.IsEmpty())
	for i := 0; i < 10; i++ {
		list.Add(T(i))
		assert.False(t, list.IsEmpty())
		list.Remove()
		assert.True(t, list.IsEmpty())
	}
	assert.True(t, list.IsEmpty())
}

func TestLen(t *testing.T) {
	list := NewLinkList[T]()
	assert.True(t, list.Len() == 0)
	num := 100
	for i := 0; i < num; {
		i++
		list.Add(T(i))
		assert.True(t, list.Len() == i)
	}
	for i := num - 1; i >= 0; i-- {
		list.Remove()
		assert.True(t, list.Len() == i)
	}
	assert.True(t, list.Len() == 0)
}

func TestSlice(t *testing.T) {
	list := NewLinkList[T]()
	num := 100
	slice := make([]T, 0, num)
	for i := 0; i < num; i++ {
		list.Add(T(i))
		slice = append(slice, T(i))
	}
	slice2 := list.Slice()
	for i := 0; i < num; i++ {
		assert.True(t, slice[i] == slice2[i])
	}

	slice3 := list.Slice(OptReverse())
	for i := num - 1; i >= 0; i-- {
		assert.True(t, slice[i] == slice3[num-i-1])
	}
}

func TestForEach(t *testing.T) {
	list := NewLinkList[T]()
	num := 100
	slice := make([]T, 0, num)
	for i := 0; i < num; i++ {
		list.Add(T(i))
		slice = append(slice, T(i))
	}
	i := 0
	list.ForEach(func(data T) bool {
		assert.True(t, slice[i] == data)
		i++
		return true
	})

	list.ForEach(func(data T) bool {
		i--
		assert.True(t, slice[i] == data)
		return true
	}, OptReverse())
}

func TestInsert(t *testing.T) {
	list := NewLinkList[T]()
	list.Insert(T(2), 0)
	list.Insert(T(1), 1)
	list.Insert(T(10), 0)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 2, 1}))

	list.Insert(T(3), 1)
	list.Insert(T(6), 1)
	list.Insert(T(5), 2)
	list.Insert(T(8), 1)
	list.Insert(T(4), 4)
	list.Insert(T(9), 1)
	list.Insert(T(7), 3)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func TestDelete(t *testing.T) {
	list := NewLinkList[T]()
	list.Insert(T(2), 0)
	list.Insert(T(1), 1)
	list.Insert(T(10), 0)
	list.Insert(T(3), 1)
	list.Insert(T(6), 1)
	list.Insert(T(5), 2)
	list.Insert(T(8), 1)
	list.Insert(T(4), 4)
	list.Insert(T(9), 1)
	list.Insert(T(7), 3)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))

	list.Delete(1)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 8, 7, 6, 5, 4, 3, 2, 1}))

	list.Delete(list.Len() - 1)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 8, 7, 6, 5, 4, 3, 2}))

	list.Delete(1)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 7, 6, 5, 4, 3, 2}))

	list.Delete(6)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 7, 6, 5, 4, 3}))

	list.Delete(4)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 7, 6, 5, 3}))

	list.Delete(2)
	assert.True(t, reflect.DeepEqual(list.Slice(), []T{10, 7, 5, 3}))
}
