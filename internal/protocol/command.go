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
)

type Command struct {
	Type CommandType

	// Common fields (used depending on command)
	Partition string
	Key       string
	Value     string

	// Control-plane
	ValueType   string
	PersistMode string
}
