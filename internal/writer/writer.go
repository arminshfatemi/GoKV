package writer

import (
	"GoKV/internal/executor"
	"bufio"
	"fmt"
)

func WriteResult(w *bufio.Writer, res executor.ExecutionResult) error {
	switch res.Type {
	case executor.ResultString:
		_, err := fmt.Fprintf(w, "+%s\n", res.Value.(string))
		return err

	case executor.ResultInt:
		_, err := fmt.Fprintf(w, ":%d\n", res.Value.(int64))
		return err

	case executor.ResultError:
		_, err := fmt.Fprintf(w, "-ERR %s\n", res.Value.(string))
		return err

	case executor.ResultArray:
		arr := res.Value.([]string)
		// first line: array length
		_, err := fmt.Fprintf(w, "*%d\n", len(arr))
		if err != nil {
			return err
		}

		for _, s := range arr {
			_, err := fmt.Fprintf(w, "$%d\n%s\n", len(s), s)
			if err != nil {
				return err
			}
		}
		return nil

	case executor.ResultNull:
		_, err := w.WriteString("$-1\r\n")
		return err
	default:
		_, err := w.WriteString("-ERR unknown result\n")
		return err
	}
}
