package protocol

type CommandType uint8

const (
	CmdUnknown CommandType = iota

	// Partition handling

	CmdCreatePartition
	CmdDropPartition
	CmdListPartitions
	CmdDescribePartition

	// Partitions operations

	CmdSet
	CmdGet
	CmdDel
	CmdIncr
	CmdStatsPartition
	CmdExists

	// Auth

	CmdAuthentication
)

type Command struct {
	PartitionKey string
	Type         CommandType
	Partition    []byte
	Args         [][]byte
}
