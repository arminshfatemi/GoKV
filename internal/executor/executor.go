package executor

import (
	"GoKV/internal/commands"
)

type Executor struct {
	command *commands.Command
	handler CommandHandler
}

func NewExecutor(command *commands.Command) (*Executor, error) {
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
