package executor

import "errors"

var (
	ErrCommandHandlerNotFound = errors.New("error couldn't find the command handler")
)
