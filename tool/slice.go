package tool

import (
	"math/rand"
	"sort"

	"github.com/grin-ch/grin-utils/math"
)

type number math.Number
type comparator math.Comparator

type opt struct {
	desc    bool
	inPlace bool
}

func newOpt(opts ...option) *opt {
	o := new(opt)
	for _, fn := range opts {
		fn(o)
	}
	return o
}

type option func(*opt)

func OptDesc() option {
	return func(o *opt) {
		o.desc = true
	}
}

func OptInPlace() option {
	return func(o *opt) {
		o.inPlace = true
	}
}

// Sort
func Sort[T comparator](arr []T, opts ...option) {
	o := newOpt(opts...)
	sliceSort(arr, o.desc)
}

func sliceSort[T comparator](arr []T, desc bool) {
	fn := func(i, j int) bool { return (arr[i] >= arr[j] && desc) || (arr[i] <= arr[j] && !desc) }
	sort.SliceStable(arr, fn)
}

// Top 获取最大/小的元素, 没有则返回默认值
func Top[T comparator](arr []T, opts ...option) (item T) {
	if len(arr) == 0 {
		return
	}

	o := newOpt(opts...)
	item = arr[0]
	var fn func(a, b T) bool
	if o.desc {
		fn = func(a, b T) bool { return a < b }
	} else {
		fn = func(a, b T) bool { return b < a }
	}
	for i := 1; i < len(arr); i++ {
		if fn(item, arr[i]) {
			item = arr[i]
		}
	}
	return
}

// Tops 获取最大/小的n个元素
func Tops[T comparator](n int, arr []T, opts ...option) []T {
	o := newOpt(opts...)
	if len(arr) < n {
		n = len(arr)
	}
	arr2 := arr
	if !o.inPlace {
		arr2 = make([]T, len(arr))
		copy(arr2, arr)
	}
	sliceSort(arr2, o.desc)
	return arr2[:n]
}

// Shuffle
func Shuffle[T any](arr []T) {
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
}

// BinarySearch 二分查找法
// 传入的arr需要各个元素升序且唯一
func BinarySearch[T number](arr []T, target T) int {
	l, r := 0, len(arr)
	for l < r {
		i := (r-l)/2 + l
		if arr[i] < target {
			l = i + 1
		} else if arr[i] > target {
			r = i
		} else {
			return i
		}
	}
	return -1
}
