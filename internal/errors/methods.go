package errors

func ErrorEmptyRequired() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 0, ErrorMessage: "Username or password fields cannot be empty."}
}

func ErrorUsernameLetter() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 1, ErrorMessage: "Username must start with a letter."}
}

func ErrorUsernameUsed() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 2, ErrorMessage: "Cannot create user. Entered username already in use."}
}

func ErrorUsernameSpaces() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 3, ErrorMessage: "Entered username contain spaces."}
}

func ErrorPasswordSpaces() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 4, ErrorMessage: "Entered password contain spaces."}
}

func ErrorNoUser() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 5, ErrorMessage: "User not found."}
}

func ErrorNoUsers() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 6, ErrorMessage: "Users not found."}
}

func ErrorUsernameInvalid() CustomError {
	return CustomError{StatusCode: 400, ErrorCode: 6, ErrorMessage: "Username is invalid."}
}
