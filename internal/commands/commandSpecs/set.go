package commandSpecs

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
)

type SetSpec struct{}

func (s *SetSpec) Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error) {
	if len(args) != 3 {
		return commands.Command{}, commands.InvalidArgsErr
	}

	if ctx.Role != auth.RoleOperator && ctx.Role != auth.RoleAdmin {
		return commands.Command{}, commands.PermissionErr
	}

	return commands.Command{
		PartitionKey: string(args[0]),
		Type:         commands.CmdSet,
		Args:         args[1:],
	}, nil
}
