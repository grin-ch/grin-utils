package time

import (
	"testing"
	"time"
)

func TestZeroClock(t *testing.T) {
	now := time.Date(2022, 11, 15, 20, 0, 0, 0, time.Local)
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				now: now,
			},
			want: "2022-11-15 20:00:00",
		},
		{
			name: "test2",
			args: args{
				now: now.AddDate(0, 0, 1),
			},
			want: "2022-11-16 20:00:00",
		},
		{
			name: "test3",
			args: args{
				now: now.AddDate(0, 0, 2),
			},
			want: "2022-11-17 20:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZeroClock(tt.args.now).Format(YYYYmmDDhhMMss)
			if got != tt.want {
				t.Errorf("ZeroClock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZeroTime(t *testing.T) {
	now := time.Date(2022, 11, 15, 20, 0, 0, 0, time.Local)
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				now: now,
			},
			want: "2022-11-15 00:00:00",
		},
		{
			name: "test2",
			args: args{
				now: now.AddDate(0, 0, 1),
			},
			want: "2022-11-16 00:00:00",
		},
		{
			name: "test3",
			args: args{
				now: now.AddDate(0, 0, 2),
			},
			want: "2022-11-17 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ZeroTime(tt.args.now).Format(YYYYmmDDhhMMss)
			if got != tt.want {
				t.Errorf("ZeroClock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstDateOfMonth(t *testing.T) {
	now := time.Date(2022, 11, 15, 20, 0, 0, 0, time.Local)
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				now: now,
			},
			want: "2022-11-01 00:00:00",
		},
		{
			name: "test2",
			args: args{
				now: now.AddDate(0, 1, 0),
			},
			want: "2022-12-01 00:00:00",
		},
		{
			name: "test3",
			args: args{
				now: now.AddDate(0, -1, 0),
			},
			want: "2022-10-01 00:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstDateOfMonth(tt.args.now).Format(YYYYmmDDhhMMss)
			if got != tt.want {
				t.Errorf("FirstDateOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastDateOfMonth(t *testing.T) {
	now := time.Date(2022, 11, 15, 20, 0, 0, 0, time.Local)
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				now: now,
			},
			want: "2022-11-30",
		},
		{
			name: "test2",
			args: args{
				now: now.AddDate(0, 1, 0),
			},
			want: "2022-12-31",
		},
		{
			name: "test3",
			args: args{
				now: now.AddDate(0, -1, 0),
			},
			want: "2022-10-31",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastDateOfMonth(tt.args.now).Format(ISODate)
			if got != tt.want {
				t.Errorf("LastDateOfMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstDateOfWeek(t *testing.T) {
	now := time.Date(2022, 11, 15, 20, 0, 0, 0, time.Local)
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				now: now,
			},
			want: "2022-11-13",
		},
		{
			name: "test2",
			args: args{
				now: now.AddDate(0, 1, 0),
			},
			want: "2022-12-11",
		},
		{
			name: "test3",
			args: args{
				now: now.AddDate(0, -1, 0),
			},
			want: "2022-10-09",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FirstDateOfWeek(tt.args.now).Format(ISODate)
			if got != tt.want {
				t.Errorf("FirstDateOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastDateOfWeek(t *testing.T) {
	now := time.Date(2022, 11, 15, 20, 0, 0, 0, time.Local)
	type args struct {
		now time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				now: now,
			},
			want: "2022-11-19",
		},
		{
			name: "test2",
			args: args{
				now: now.AddDate(0, 1, 0),
			},
			want: "2022-12-17",
		},
		{
			name: "test3",
			args: args{
				now: now.AddDate(0, -1, 0),
			},
			want: "2022-10-15",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LastDateOfWeek(tt.args.now).Format(ISODate)
			if got != tt.want {
				t.Errorf("LastDateOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
