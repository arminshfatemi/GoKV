package protocol

import (
	"GoKV/internal/commands"
	"GoKV/internal/commands/commandSpecs"
)

type CommandSpec interface {
	Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error)
}

type Node struct {
	Key      []byte
	Children []*Node
	Spec     CommandSpec
}

type MatchResult struct {
	Spec CommandSpec
	Args [][]byte
}

var rootNode = &Node{
	Key:  nil,
	Spec: nil,
	Children: []*Node{
		{
			Key:  []byte("PARTITION"),
			Spec: nil,
			Children: []*Node{
				{
					Key:      []byte("CREATE"),
					Spec:     &commandSpecs.CreatePartitionSpec{},
					Children: nil,
				},
				{
					Key:      []byte("DROP"),
					Spec:     &commandSpecs.DropPartitionSpec{},
					Children: nil,
				},
				{
					Key:      []byte("DESCRIBE"),
					Spec:     &commandSpecs.DescribePartitionSpec{},
					Children: nil,
				},
				{
					Key:      []byte("STATS"),
					Spec:     &commandSpecs.StatsPartitionSpec{},
					Children: nil,
				},
				{
					Key:      []byte("LIST"),
					Spec:     &commandSpecs.ListPartitionSpec{},
					Children: nil,
				},
			},
		},
		{
			Key:      []byte("GET"),
			Spec:     &commandSpecs.GetSpec{},
			Children: nil,
		},
		{
			Key:      []byte("SET"),
			Spec:     &commandSpecs.SetSpec{},
			Children: nil,
		},
		{
			Key:      []byte("INCR"),
			Spec:     &commandSpecs.IncrSpec{},
			Children: nil,
		},
		{
			Key:      []byte("DEL"),
			Spec:     &commandSpecs.DelSpec{},
			Children: nil,
		},
		{
			Key:      []byte("EXISTS"),
			Spec:     &commandSpecs.ExistsSpec{},
			Children: nil,
		},
	},
}
