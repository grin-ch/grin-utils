package tool

import (
	"math"
	"testing"
)

func TestInitSnowflakeNode(t *testing.T) {
	type args struct {
		node int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				node: math.MaxInt64,
			},
			wantErr: true,
		},
		{
			name: "test2",
			args: args{
				node: 1000,
			},
			wantErr: false,
		},
		{
			name: "test3",
			args: args{
				node: -1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitSnowflakeNode(tt.args.node); (err != nil) != tt.wantErr {
				t.Errorf("InitSnowflakeNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
func BenchmarkNewSnowFlakeID(b *testing.B) {
	InitSnowflakeNode(0)
	for n := 0; n < b.N; n++ {
		NewSnowFlakeID()
	}
}

func BenchmarkNewSnowFlakeID2(b *testing.B) {
	InitSnowflakeNode(0)
	for n := 0; n < b.N; n++ {
		nextSnowflakeID()
	}
}

func BenchmarkNewSnowFlakeID3(b *testing.B) {
	node, _ := NewSnowFlakeNode(0)
	for n := 0; n < b.N; n++ {
		node.Generate()
	}
}
