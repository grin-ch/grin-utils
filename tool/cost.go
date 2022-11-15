package tool

import "time"

// 计算方法耗时,返回一个纳秒时间
type callCost func() int64

// defer Cost()()
func Cost() callCost {
	now := time.Now()
	return func() int64 {
		return -int64(time.Until(now))
	}
}
