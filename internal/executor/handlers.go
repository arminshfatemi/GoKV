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
