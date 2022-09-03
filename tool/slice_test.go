package tool

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	arr := []int{1, -3, 45, 5, 14, 1, 3, 123, 123, 543, 5, 6, 345, 234, 6}
	descArr := []int{543, 345, 234, 123, 123, 45, 14, 6, 6, 5, 5, 3, 1, 1, -3}
	ascArr := []int{-3, 1, 1, 3, 5, 5, 6, 6, 14, 45, 123, 123, 234, 345, 543}
	Sort(arr, OptDesc())
	for i := 0; i < len(arr); i++ {
		if arr[i] != descArr[i] {
			t.Fatalf("arr[%d]:%d, descArr[%d]:%d", i, arr[i], i, descArr[i])
		}
	}
	Sort(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != ascArr[i] {
			t.Fatalf("arr[%d]:%d, ascArr[%d]:%d", i, arr[i], i, ascArr[i])
		}
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []T
		target T
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 0,
			},
			want: -1,
		},
		{
			name: "test2",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 1,
			},
			want: 0,
		},
		{
			name: "test3",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "test4",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 3,
			},
			want: 2,
		},
		{
			name: "test5",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 4,
			},
			want: 3,
		},
		{
			name: "test6",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 5,
			},
			want: 4,
		},
		{
			name: "test7",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 6,
			},
			want: 5,
		},
		{
			name: "test8",
			args: args{
				arr:    []T{1, 2, 3, 4, 5, 6},
				target: 7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTop(t *testing.T) {
	arr := []int{1, -3, 45, 5, 14, 1, 3, 123, 123, 543, 5, 6, 345, 234, 6}
	top := Top(arr, OptDesc())
	if top != 543 {
		t.Fatalf("top:%d", top)
	} else {
		t.Logf("correct top:%d", top)
	}

	top2 := Top(arr)
	if top2 != -3 {
		t.Fatalf("top2:%d", top2)
	} else {
		t.Logf("correct top2:%d", top2)
	}
}

func TestTops(t *testing.T) {
	type args struct {
		n    int
		arr  []T
		opts []option
	}
	tests := []struct {
		name      string
		args      args
		want      []T
		isInPlace bool
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				n: 0,
			},
			want: []T{},
		},
		{
			name: "test1",
			args: args{
				n:    1,
				arr:  []T{1, 3, 5, 2, 4, 6},
				opts: []option{OptInPlace()},
			},
			want:      []T{1},
			isInPlace: true,
		},
		{
			name: "test2",
			args: args{
				n:   1,
				arr: []T{1, 3, 5, 2, 4, 6},
			},
			want: []T{1},
		},
		{
			name: "test3",
			args: args{
				n:    3,
				arr:  []T{1, 3, 5, 2, 4, 6},
				opts: []option{OptInPlace(), OptDesc()},
			},
			want:      []T{6, 5, 4},
			isInPlace: true,
		},
		{
			name: "test",
			args: args{
				n:    3,
				arr:  []T{1, 3, 5, 2, 4, 6},
				opts: []option{},
			},
			want: []T{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tops(tt.args.n, tt.args.arr, tt.args.opts...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tops() = %v, want %v", got, tt.want)
			}
			if tt.isInPlace {
				for i := 0; i < len(got); i++ {
					if got[i] != tt.args.arr[i] {
						t.Errorf("got[%d] = %v, want %v", i, got[i], tt.args.arr[i])
					}
				}
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	rand.Seed(time.Now().Unix())
	type args struct {
		arr []T
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				arr: []T{},
			},
		},
		{
			name: "test2",
			args: args{
				arr: []T{1},
			},
		},
		{
			name: "test3",
			args: args{
				arr: []T{1, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Shuffle(tt.args.arr)
			fmt.Println(tt.args.arr)
		})
	}
}
