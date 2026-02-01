package tcp

import (
	"GoKV/internal/executor"
	"GoKV/internal/protocol"
	writer2 "GoKV/internal/writer"
	"bufio"
	"bytes"
	"net"
)

const (
	maxLineSize = 64 * 1024 // 64KB per command
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReaderSize(conn, maxLineSize)
	writer := bufio.NewWriter(conn)
	parser := protocol.NewParser()

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			return
		}

		line = bytes.TrimSpace(line)

		cmd, err := parser.Parse(line)
		if err != nil {
			writer.WriteString("-ERR " + err.Error() + "\n")
			writer.Flush()
			continue
		}

		e, err := executor.NewExecutor(cmd)
		if err != nil {
			writer.WriteString("-ERR " + err.Error() + "\n")
		}

		result, err := e.Execute()
		if err != nil {
			writer.WriteString("-ERR " + err.Error() + "\n")
		}

		if err := writer2.WriteResult(writer, result); err != nil {
			panic(e)
		}
		writer.Flush()
	}
}
