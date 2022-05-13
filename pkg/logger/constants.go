package logger

const (
	ErrLevelInfo       = "info"
	ErrLevelWarning    = "warning"
	ErrLevelError      = "error"
	ErrLevelFatal      = "fatal"
	TypeUser           = "*user.User"
	TypeUsersMap       = "map[int]*user.User"
	GotUserData        = "Got user's data."
	GotAllUsersData    = "Got all-user's data."
	TypeUserperm       = "*userperm.Userperm"
	TypeUserpermsMap   = "map[int]*userperm.Userperm"
	GotUserPermissions = "Got user's permissions."
)
