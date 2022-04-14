package vh

type customError struct {
	StatusCode   int
	ErrorCode    int
	ErrorMessage string
}

func ErrorEmptyRequired() customError {
	return customError{StatusCode: 400, ErrorCode: 0, ErrorMessage: "Username or password fields cannot be empty."}
}

func ErrorUsernameLetter() customError {
	return customError{StatusCode: 400, ErrorCode: 1, ErrorMessage: "Username must start with a letter."}
}

func ErrorUsernameUsed() customError {
	return customError{StatusCode: 400, ErrorCode: 2, ErrorMessage: "Cannot create user. Entered username already in use."}
}

func ErrorUsernameSpaces() customError {
	return customError{StatusCode: 400, ErrorCode: 3, ErrorMessage: "Entered username contain spaces."}
}

func ErrorPasswordSpaces() customError {
	return customError{StatusCode: 400, ErrorCode: 4, ErrorMessage: "Entered password contain spaces."}
}

func ErrorNoUser() customError {
	return customError{StatusCode: 400, ErrorCode: 5, ErrorMessage: "User not found."}
}

func ErrorNoUsers() customError {
	return customError{StatusCode: 400, ErrorCode: 6, ErrorMessage: "Users not found."}
}

func ErrorUsernameInvalid() customError {
	return customError{StatusCode: 400, ErrorCode: 6, ErrorMessage: "Username is invalid."}
}
