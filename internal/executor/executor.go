package executor

import "GoKV/internal/protocol"

type Executor struct {
	command *protocol.Command
	handler CommandHandler
}

func NewExecutor(command *protocol.Command) (*Executor, error) {
	// find the handler for this command
	h, ok := CommandTable[command.Type]
	if !ok {
		return nil, ErrCommandHandlerNotFound
	}

	// create Executor
	e := &Executor{
		command: command,
		handler: h,
	}
	return e, nil
}

func (e *Executor) Execute() ExecutionResult {
	return e.handler(e.command)
}
