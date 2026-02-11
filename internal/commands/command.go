package commands

type CommandType uint8

const (
	CmdUnknown CommandType = iota

	// Partition handling

	CmdCreatePartition
	CmdDropPartition
	CmdListPartitions
	CmdDescribePartition
	CmdStatsPartition

	// Partitions operations

	CmdSet
	CmdGet
	CmdDel
	CmdIncr
	CmdExists

	// Auth

	CmdAuthentication
)

type Command struct {
	PartitionKey string
	Type         CommandType
	Args         [][]byte
}
