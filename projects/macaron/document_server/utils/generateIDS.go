package utils

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

// todo: 使用雪花算法生成id,为int64
func GenerateIDS() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	id := node.Generate()
	return id.Int64()
}
