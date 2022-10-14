package safe

import (
	"fmt"
	"sync"
	"testing"
)

func TestGo(t *testing.T) {
	type args struct {
		fn   func()
		opts []Option
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{
				fn: func() {
					fmt.Print("hello t1")
				},
				opts: []Option{OptCallback(func() { fmt.Println("world!") })},
			},
		},
		{
			name: "t2",
			args: args{
				fn: func() {
					fmt.Print("hello t2")
				},
			},
		},
		{
			name: "t3",
			args: args{
				fn: func() {
					panic("i am bad")
				},
				opts: []Option{OptLog(func(err error) { fmt.Println("panic err:", err.Error()) })},
			},
		},
		{
			name: "t4",
			args: args{
				fn: func() {
					panic("i am bad")
				},
			},
		},
		{
			name: "t5",
			args: args{
				fn: func() {
					fmt.Print("hello ")
				},
				opts: []Option{OptLog(func(err error) { fmt.Println("panic err:", err.Error()) })},
			},
		},
		{
			name: "t6",
			args: args{
				fn: func() {
					panic("hello")
				},
				opts: []Option{
					OptCallback(func() { fmt.Println("world!") }),
					OptFinal(func() { panic("finaly panic") }),
				},
			},
		},
		{
			name: "t7",
			args: args{
				fn: func() {
					fmt.Print("hello ")
				},
				opts: []Option{
					OptCallback(func() { fmt.Println("world!") }),
					OptLog(func(err error) { fmt.Println("panic err:", err.Error()) }),
					OptFinal(func() { panic("finaly") }),
				},
			},
		},
		{
			name: "t8",
			args: args{
				fn: func() {
					panic("hello")
				},
				opts: []Option{
					OptCallback(func() { fmt.Println("world!") }),
					OptLog(func(err error) { fmt.Println("panic err:", err.Error()) }),
					OptFinal(func() { fmt.Println("finaly") }),
				},
			},
		},
		{
			name: "t9",
			args: args{
				fn: func() {
					panic(" panic ")
				},
				opts: []Option{
					OptLog(func(err error) { panic(" err handle panic ") }),
				},
			},
		},
	}
	wait := sync.WaitGroup{}
	for _, tt := range tests {
		wait.Add(1)
		t.Run(tt.name, func(t *testing.T) {
			Go(tt.args.fn, tt.args.opts...)
			wait.Done()
		})
	}
	wait.Wait()
	fmt.Println("-----------------")
	fmt.Println("go on")
}
