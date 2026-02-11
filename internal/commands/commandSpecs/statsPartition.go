package commandSpecs

import (
	"GoKV/internal/auth"
	"GoKV/internal/commands"
)

type StatsPartitionSpec struct{}

func (s *StatsPartitionSpec) Build(args [][]byte, ctx commands.BuildContext) (commands.Command, error) {
	if len(args) != 1 {
		return commands.Command{}, commands.InvalidArgsErr
	}

	if ctx.Role != auth.RoleAdmin {
		return commands.Command{}, commands.PermissionErr
	}

	return commands.Command{
		PartitionKey: string(args[0]),
		Type:         commands.CmdStatsPartition,
		Args:         nil,
	}, nil

}
