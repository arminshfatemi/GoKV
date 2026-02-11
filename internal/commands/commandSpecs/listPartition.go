package commandSpecs

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
)

type ListPartitionSpec struct{}

func (l *ListPartitionSpec) Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error) {
	if len(args) != 0 {
		return commands.Command{}, commands.InvalidArgsErr
	}

	if ctx.Role != auth.RoleAdmin {
		return commands.Command{}, commands.PermissionErr
	}

	return commands.Command{
		PartitionKey: "",
		Type:         commands.CmdListPartitions,
		Args:         nil,
	}, nil
}
