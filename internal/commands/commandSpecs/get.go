package commandSpecs

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
)

type GetSpec struct{}

func (l *GetSpec) Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error) {
	if len(args) != 2 {
		return commands.Command{}, commands.InvalidArgsErr
	}

	if ctx.Role != auth.RoleReader && ctx.Role != auth.RoleAdmin {
		return commands.Command{}, commands.PermissionErr
	}

	return commands.Command{
		PartitionKey: string(args[0]),
		Type:         commands.CmdGet,
		Args:         args[1:],
	}, nil
}
