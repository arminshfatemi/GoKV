package commands

import "errors"

var (
	InvalidArgsErr = errors.New("invalid args")
	PermissionErr  = errors.New("permission denied")
)
