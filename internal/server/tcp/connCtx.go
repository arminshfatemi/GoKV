package tcp

import "GoKV/internal/auth"

type ConnCtx struct {
	User   *auth.User
	Authed bool
}
