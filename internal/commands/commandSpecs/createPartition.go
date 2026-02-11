package commandSpecs

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
)

type CreatePartitionSpec struct{}

func (c *CreatePartitionSpec) Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error) {
	if len(args) != 3 {
		return commands.Command{}, commands.InvalidArgsErr
	}

	if ctx.Role != auth.RoleAdmin {
		return commands.Command{}, commands.PermissionErr
	}

	return commands.Command{
		PartitionKey: string(args[0]),
		Type:         commands.CmdCreatePartition,
		Args:         args[1:],
	}, nil
}
