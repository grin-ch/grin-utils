package time

import (
	"time"
)

func ZeroClock(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
}

func ZeroTime(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
}

// 月初第一天
func FirstDateOfMonth(now time.Time) time.Time {
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func LastDateOfMonth(now time.Time) time.Time {
	return FirstDateOfMonth(now).AddDate(0, 1, -1)
}

// 规定周日为一周的第一天
func FirstDateOfWeek(now time.Time) time.Time {
	diff := time.Sunday - now.Weekday()
	return now.AddDate(0, 0, int(diff))
}

func LastDateOfWeek(now time.Time) time.Time {
	diff := time.Saturday - now.Weekday()
	return now.AddDate(0, 0, int(diff))
}
