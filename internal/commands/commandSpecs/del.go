package commandSpecs

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
)

type DelSpec struct{}

func (l *DelSpec) Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error) {
	if len(args) < 2 {
		return commands.Command{}, commands.InvalidArgsErr
	}

	if ctx.Role != auth.RoleAdmin && ctx.Role != auth.RoleOperator {
		return commands.Command{}, commands.PermissionErr
	}

	return commands.Command{
		PartitionKey: string(args[0]),
		Type:         commands.CmdDel,
		Args:         args[1:],
	}, nil
}
