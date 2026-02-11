package commandSpecs

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
)

type ExistsSpec struct{}

func (l *ExistsSpec) Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error) {
	if len(args) != 2 {
		return commands.Command{}, commands.InvalidArgsErr
	}

	if ctx.Role != auth.RoleAdmin && ctx.Role != auth.RoleReader {
		return commands.Command{}, commands.PermissionErr
	}

	return commands.Command{
		PartitionKey: string(args[0]),
		Type:         commands.CmdExists,
		Args:         args[1:],
	}, nil
}
