package protocol

type CommandType uint8

const (
	CmdUnknown CommandType = iota
	CmdCreatePartition
	CmdDropPartition
	CmdListPartitions
	CmdDescribePartition
	CmdSet
	CmdGet
	CmdDel
	CmdIncr
	CmdStatsPartition
	CmdExists
)

type Command struct {
	PartitionKey string
	Type         CommandType
	Partition    []byte
	Args         [][]byte
}
