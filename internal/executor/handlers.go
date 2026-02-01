package executor

import (
	"GoKV/internal/partitions"
	"GoKV/internal/protocol"
)

type CommandHandler func(command *protocol.Command) (ExecutionResult, error)

var CommandTable = map[protocol.CommandType]CommandHandler{}

func init() {
	addHandler(protocol.CmdCreatePartition, createPartitionHandler)
	addHandler(protocol.CmdListPartitions, listPartitionHandler)
	addHandler(protocol.CmdDropPartition, dropPartitionHandler)

	addHandler(protocol.CmdDel, delHandler)
	addHandler(protocol.CmdGet, getHandler)
	addHandler(protocol.CmdSet, setHandler)
	addHandler(protocol.CmdIncr, incrHandler)
}

func addHandler(commandType protocol.CommandType, handler CommandHandler) {
	CommandTable[commandType] = handler
}

func createPartitionHandler(command *protocol.Command) (ExecutionResult, error) {
	if command.Type != protocol.CmdCreatePartition {
		panic("createPartitionHandler called with wrong command type")
	}

	vt, err := partitions.ParseValueType(command.ValueType)
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}, nil
	}

	pm, err := partitions.ParsePersistMode(command.PersistMode)
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}, nil
	}

	// create a config
	cfg := partitions.PartitionConfig{
		Name:        command.Partition,
		Schema:      vt,
		PersistMode: pm,
	}

	// create partition
	if err := partitions.CreatePartition(cfg); err != nil {
		return ExecutionResult{Type: ResultError, Value: err.Error()}, nil
	}

	// create result
	r := ExecutionResult{
		Type:  ResultString,
		Value: "OK",
	}

	return r, nil
}

func listPartitionHandler(command *protocol.Command) (ExecutionResult, error) {
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
	return r, nil
}

func dropPartitionHandler(command *protocol.Command) (ExecutionResult, error) {
	if command.Type != protocol.CmdDropPartition {
		panic("dropPartitionHandler called with wrong command type")
	}

	// list the partitions
	err := partitions.DropPartition(command.Partition)
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}, nil
	}

	// create result
	r := ExecutionResult{
		Type:  ResultString,
		Value: "OK",
	}
	return r, nil
}

func getHandler(command *protocol.Command) (ExecutionResult, error) {
	partitionName := command.Partition
	key := command.Key

	// get partition
	p, ok := partitions.GetPartition(partitionName)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}, nil
	}

	v, ok := p.Get(key)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrKeyNotFound.Error(),
		}, nil
	}

	rType := ResultString
	if p.Schema == partitions.INT {
		rType = ResultInt
	}

	r := ExecutionResult{
		Type:  rType,
		Value: v,
	}

	return r, nil
}

func setHandler(command *protocol.Command) (ExecutionResult, error) {
	partitionName := command.Partition
	key := command.Key
	value := command.Value

	p, ok := partitions.GetPartition(partitionName)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}, nil
	}

	if err := p.Set(key, []byte(value)); err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}, nil
	}

	r := ExecutionResult{
		Type:  ResultString,
		Value: "OK",
	}
	return r, nil
}

func delHandler(command *protocol.Command) (ExecutionResult, error) {
	partitionName := command.Partition
	key := command.Key

	p, ok := partitions.GetPartition(partitionName)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}, nil
	}

	deleted := p.Del(key)

	var count int64
	if deleted {
		count = 1
	}

	return ExecutionResult{
		Type:  ResultInt,
		Value: count,
	}, nil
}

func incrHandler(command *protocol.Command) (ExecutionResult, error) {
	partitionName := command.Partition
	key := command.Key

	p, ok := partitions.GetPartition(partitionName)
	if !ok {
		return ExecutionResult{
			Type:  ResultError,
			Value: partitions.ErrPartitionNotFound.Error(),
		}, nil
	}

	v, err := p.Incr(key)
	if err != nil {
		return ExecutionResult{
			Type:  ResultError,
			Value: err.Error(),
		}, err
	}

	r := ExecutionResult{
		Type:  ResultInt,
		Value: v,
	}

	return r, nil
}
