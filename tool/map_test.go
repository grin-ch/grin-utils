package tool

import (
	"testing"
)

type K string
type V int
type T int

func TestMapFlip(t *testing.T) {
	type args struct {
		kvs map[K]V
		fn  CustomKey[V, T]
	}
	tests := []struct {
		name string
		args args
		size int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				fn: func(i V) T { return T(i) },
			},
			size: 3,
		},
		{
			name: "test2",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 1,
				},
				fn: func(i V) T { return T(i) },
			},
			size: 1,
		},
		{
			name: "test2",
			args: args{
				fn: func(i V) T { return T(i) },
			},
			size: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapFlip(tt.args.kvs, tt.args.fn); len(got) != tt.size {
				t.Errorf("len(got) = %d,  want size = %d ", len(got), tt.size)
			}
		})
	}
}

func TestMapFlipSlice(t *testing.T) {
	type args struct {
		kvs map[K]V
		fn  CustomKey[V, T]
	}
	tests := []struct {
		name string
		args args
		size int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 2,
					"c": 3,
				},
				fn: func(i V) T { return T(i) },
			},
			size: 3,
		},
		{
			name: "test2",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 1,
				},
				fn: func(i V) T { return T(i) },
			},
			size: 1,
		},
		{
			name: "test2",
			args: args{
				fn: func(i V) T { return T(i) },
			},
			size: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapFlipSlice(tt.args.kvs, tt.args.fn); len(got) != tt.size {
				t.Errorf("len(got) = %d,  want size = %d ", len(got), tt.size)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	type args struct {
		kvs map[K]V
	}
	tests := []struct {
		name string
		args args
		size int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			size: 3,
		},
		{
			name: "test2",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 1,
				},
			},
			size: 2,
		},
		{
			name: "test3",
			args: args{},
			size: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				if got := MapKeys(tt.args.kvs); len(got) != tt.size {
					t.Errorf("len(got) = %d,  want size = %d ", len(got), tt.size)
				}
			})
		})
	}
}

func TestMapValues(t *testing.T) {
	type args struct {
		kvs map[K]V
	}
	tests := []struct {
		name string
		args args
		size int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 2,
					"c": 3,
				},
			},
			size: 3,
		},
		{
			name: "test2",
			args: args{
				kvs: map[K]V{
					"a": 1,
					"b": 1,
				},
			},
			size: 2,
		},
		{
			name: "test3",
			args: args{},
			size: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				if got := MapValues(tt.args.kvs); len(got) != tt.size {
					t.Errorf("len(got) = %d,  want size = %d ", len(got), tt.size)
				}
			})
		})
	}
}
