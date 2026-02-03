package executor

import (
	"GoKV/internal/partitions"
	"GoKV/internal/protocol"
)

type CommandHandler func(command *protocol.Command) ExecutionResult

var CommandTable = map[protocol.CommandType]CommandHandler{}

func init() {
	addHandler(protocol.CmdCreatePartition, createPartitionHandler)
	addHandler(protocol.CmdListPartitions, listPartitionHandler)
	addHandler(protocol.CmdDropPartition, dropPartitionHandler)
	addHandler(protocol.CmdDescribePartition, describeHandler)
	addHandler(protocol.CmdStatsPartition, statsHandler)

	addHandler(protocol.CmdDel, delHandler)
	addHandler(protocol.CmdGet, getHandler)
	addHandler(protocol.CmdSet, setHandler)
	addHandler(protocol.CmdIncr, incrHandler)
	addHandler(protocol.CmdExists, existsHandler)
}

func addHandler(commandType protocol.CommandType, handler CommandHandler) {
	CommandTable[commandType] = handler
}

func createPartitionHandler(command *protocol.Command) ExecutionResult {
	if command.Type != protocol.CmdCreatePartition {
		panic("createPartitionHandler called with wrong command type")
	}

	vtRaw := command.Args[0]
	pmRaw := command.Args[1]

	vt, err := partitions.ParseValueTypeBytes(vtRaw)
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}
	}

	pm, err := partitions.ParsePersistModeBytes(pmRaw)
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}
	}

	// create a config
	cfg := partitions.PartitionConfig{
		Name:        string(command.Partition),
		Schema:      vt,
		PersistMode: pm,
	}

	// create partition
	if err := partitions.CreatePartition(cfg); err != nil {
		return ExecutionResult{Type: ResultError, Value: err.Error()}
	}

	// create result
	r := ExecutionResult{
		Type:  ResultString,
		Value: "OK",
	}

	return r
}

func listPartitionHandler(command *protocol.Command) ExecutionResult {
	if command.Type != protocol.CmdListPartitions {
		panic("listPartitionHandler called with wrong command type")
	}

	// list the partitions
	partitionsSlice := partitions.ListPartitions()

	// create result
	r := ExecutionResult{
		Type:  ResultArray,
		Value: partitionsSlice,
	}
	return r
}

func dropPartitionHandler(command *protocol.Command) ExecutionResult {
	if command.Type != protocol.CmdDropPartition {
		panic("dropPartitionHandler called with wrong command type")
	}

	// list the partitions
	err := partitions.DropPartition(string(command.Partition))
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}
	}

	// create result
	r := ExecutionResult{
		Type:  ResultString,
		Value: "OK",
	}
	return r
}

func getHandler(command *protocol.Command) ExecutionResult {
	partitionName := command.Partition
	key := string(command.Args[0])

	// get partition
	p, ok := partitions.GetPartition(string(partitionName))
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	v, ok := p.Get(key)
	if !ok {
		return ExecutionResult{
			Type: ResultNull,
		}
	}

	rType := ResultString
	if p.Schema == partitions.INT {
		rType = ResultInt
	}

	r := ExecutionResult{
		Type:  rType,
		Value: v,
	}

	return r
}

func setHandler(command *protocol.Command) ExecutionResult {
	partitionName := command.Partition
	key := string(command.Args[0])
	value := command.Args[1]

	p, ok := partitions.GetPartition(string(partitionName))
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	if err := p.Set(key, value); err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}
	}

	r := ExecutionResult{
		Type:  ResultString,
		Value: "OK",
	}
	return r
}

func delHandler(command *protocol.Command) ExecutionResult {
	partitionName := command.Partition

	keys := make([]string, 0, len(command.Args))
	for _, arg := range command.Args {
		keys = append(keys, string(arg))
	}

	p, ok := partitions.GetPartition(string(partitionName))
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	count := p.BulkDel(keys)

	return ExecutionResult{
		Type:  ResultInt,
		Value: count,
	}
}

func incrHandler(command *protocol.Command) ExecutionResult {
	partitionName := command.Partition
	key := string(command.Args[0])

	p, ok := partitions.GetPartition(string(partitionName))
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	v, err := p.Incr(key)
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}
	}

	r := ExecutionResult{
		Type:  ResultInt,
		Value: v,
	}

	return r
}

func describeHandler(command *protocol.Command) ExecutionResult {
	partitionName := command.Partition

	p, ok := partitions.GetPartition(string(partitionName))
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	v := p.Describe()

	r := ExecutionResult{
		Type:  ResultArray,
		Value: v,
	}

	return r
}

func statsHandler(command *protocol.Command) ExecutionResult {
	partitionName := command.Partition

	p, ok := partitions.GetPartition(string(partitionName))
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	stats := p.Stat()

	return ExecutionResult{
		Type:  ResultArray,
		Value: stats,
	}
}

func existsHandler(command *protocol.Command) ExecutionResult {
	partitionName := command.Partition
	key := string(command.Args[0])

	p, ok := partitions.GetPartition(string(partitionName))
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	e := p.Exists(key)

	var count int64
	if e {
		count = 1
	}

	r := ExecutionResult{
		Type:  ResultInt,
		Value: count,
	}

	return r
}
