package generator

import (
	"github.com/bwmarrin/snowflake"
	"love_knot/pkg/logger"
)

func IDGenerator() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Panic(err)
	}

	return node.Generate().Int64()
}
