package utils

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var snowflakeNodeVar *snowflake.Node

func init() {
	snowflakeNode, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Sprintf("global_id init err => %s", err)
	}

	snowflakeNodeVar = snowflakeNode
}
