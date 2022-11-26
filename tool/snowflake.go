package tool

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	snowOnce        sync.Once
	nextSnowflakeID func() int64 = func() int64 {
		return -1
	}
)

// InitSnowflakeNode 只能初始化一次
func InitSnowflakeNode(node int64) error {
	var sn *snowflake.Node
	var err error
	snowOnce.Do(func() {
		sn, err = snowflake.NewNode(node)
		nextSnowflakeID = func() int64 {
			return int64(sn.Generate())
		}
	})
	return err
}

// NewSnowFlakeID 生成下一个SnowFlakeID
// 若未成功调用过InitSnowflakeNode则返回-1
func NewSnowFlakeID() int64 {
	return nextSnowflakeID()
}

func NewSnowFlakeNode(node int) (*snowflake.Node, error) {
	return snowflake.NewNode(1)
}
