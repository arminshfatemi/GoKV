package commands

import "GoKV/internal/auth"

type BuildContext struct {
	Username string
	Role     auth.Role
}
