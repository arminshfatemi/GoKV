package tcp

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
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

func handleConnection(conn net.Conn, store *auth.Store) {
	defer conn.Close()

	reader := bufio.NewReaderSize(conn, maxLineSize)
	writer := bufio.NewWriter(conn)
	parser := protocol.NewParser()
	ctx := ConnCtx{Authed: false}

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

		if cmd.Type == commands.CmdAuthentication {
			user, ok := store.Authenticate(string(cmd.Args[0]), cmd.Args[1])
			if !ok {
				writer.WriteString("-invalid credentials\n")
				continue
			}
			ctx.User = user
			ctx.Authed = true
			writer.WriteString("+OK\n")
			writer.Flush()
			continue
		}

		e, err := executor.NewExecutor(cmd)
		if err != nil {
			writer.WriteString("-ERR " + err.Error() + "\n")
		}

		result := e.Execute()

		if err := writer2.WriteResult(writer, result); err != nil {
			panic(e)
		}
		writer.Flush()
	}
}
