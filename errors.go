package vh

var Errors = map[int]string{
	ErrorServerDebug:    "Server Debug Error",
	ErrorEmptyRequired:  "Username or password fields cannot be empty.",
	ErrorUsernameLetter: "Username must start with a letter.",
	ErrorUsernameUsed:   "Cannot create user. Entered username already in use.",
	ErrorUsernameSpaces: "Entered username contain spaces.",
	ErrorPasswordSpaces: "Entered password contain spaces.",
	ErrorNoUser:         "User not found.",
	ErrorNoUsers:        "Users not found.",
}
