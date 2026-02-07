package auth

import "strings"

type Role uint8

const (
	RoleNone Role = iota
	RoleReader
	RoleOperator
	RoleAdmin
)

const (
	StrRoleNone     = "None"
	StrRoleReader   = "Reader"
	StrRoleOperator = "Operator"
	StrRoleAdmin    = "Admin"
)

func (r Role) String() string {
	switch r {
	case RoleReader:
		return "reader"
	case RoleOperator:
		return "operator"
	case RoleAdmin:
		return "admin"
	default:
		return "none"
	}
}

func ParseRoleStr(r string) Role {
	switch strings.ToUpper(r) {
	case StrRoleReader:
		return RoleReader
	case StrRoleOperator:
		return RoleOperator
	case StrRoleAdmin:
		return RoleAdmin
	default:
		return RoleNone
	}
}
