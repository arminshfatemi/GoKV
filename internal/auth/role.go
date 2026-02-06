package auth

type Role uint8

const (
	RoleNone Role = iota
	RoleReader
	RoleOperator
	RoleAdmin
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
