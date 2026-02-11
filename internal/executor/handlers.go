package executor

import (
	"GoKV/internal/commands"
	"GoKV/internal/partitions"
)

type CommandHandler func(command *commands.Command) ExecutionResult

var CommandTable = map[commands.CommandType]CommandHandler{}

func init() {
	addHandler(commands.CmdCreatePartition, CreatePartitionHandler)
	addHandler(commands.CmdListPartitions, listPartitionHandler)
	addHandler(commands.CmdDropPartition, dropPartitionHandler)
	addHandler(commands.CmdDescribePartition, describeHandler)
	addHandler(commands.CmdStatsPartition, statsHandler)

	addHandler(commands.CmdDel, delHandler)
	addHandler(commands.CmdGet, getHandler)
	addHandler(commands.CmdSet, setHandler)
	addHandler(commands.CmdIncr, incrHandler)
	addHandler(commands.CmdExists, existsHandler)
}

func addHandler(commandType commands.CommandType, handler CommandHandler) {
	CommandTable[commandType] = handler
}

func CreatePartitionHandler(command *commands.Command) ExecutionResult {
	if command.Type != commands.CmdCreatePartition {
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
		Name:        command.PartitionKey,
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

func listPartitionHandler(command *commands.Command) ExecutionResult {
	if command.Type != commands.CmdListPartitions {
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

func dropPartitionHandler(command *commands.Command) ExecutionResult {
	if command.Type != commands.CmdDropPartition {
		panic("dropPartitionHandler called with wrong command type")
	}

	// list the partitions
	err := partitions.DropPartition(command.PartitionKey)
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

func getHandler(command *commands.Command) ExecutionResult {
	// get partition
	p, ok := partitions.GetPartition(command.PartitionKey)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	v, ok := p.Get(command.Args[0])
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

func setHandler(command *commands.Command) ExecutionResult {
	p, ok := partitions.GetPartition(command.PartitionKey)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	if err := p.Set(command.Args[0], command.Args[1]); err != nil {
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

func delHandler(command *commands.Command) ExecutionResult {
	p, ok := partitions.GetPartition(command.PartitionKey)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	count := p.BulkDel(command.Args)

	return ExecutionResult{
		Type:  ResultInt,
		Value: count,
	}
}

func incrHandler(command *commands.Command) ExecutionResult {
	p, ok := partitions.GetPartition(command.PartitionKey)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	v, err := p.Incr(command.Args[0])
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

func describeHandler(command *commands.Command) ExecutionResult {
	p, ok := partitions.GetPartition(command.PartitionKey)
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

func statsHandler(command *commands.Command) ExecutionResult {
	p, ok := partitions.GetPartition(command.PartitionKey)
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

func existsHandler(command *commands.Command) ExecutionResult {
	p, ok := partitions.GetPartition(command.PartitionKey)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}
	}

	count := p.Exists(command.Args)

	r := ExecutionResult{
		Type:  ResultInt,
		Value: count,
	}

	return r
}
