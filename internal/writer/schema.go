package writer

import (
	"GoKV/internal/executor"
	"bufio"
)

type Writer interface {
	Write(w *bufio.Writer, res executor.ExecutionResult) error
	WriteError(w *bufio.Writer, err error) error
}
