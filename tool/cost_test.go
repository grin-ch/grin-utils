package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestCostUnix(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
		},
		{
			name: "test2",
		},
		{
			name: "test3",
		},
		{
			name: "test4",
		},
		{
			name: "test5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Cost()
			defer func() {
				fmt.Println(got())
			}()
			time.Sleep(1 * time.Second)
		})
	}
}
