package util

import "github.com/bwmarrin/snowflake"

var IdWork *snowflake.Node

func init() {
	IdWork = newIdWork()
}

func newIdWork() *snowflake.Node {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	return node
}
