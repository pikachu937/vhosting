package user

const (
	ErrorCreateUser          = "Cannot create user. "
	ErrorGetUser             = "Cannot get user. "
	ErrorGetAllUsers         = "Cannot get all users. "
	ErrorPartiallyUpdateUser = "Cannot partially update user. "
	ErrorDeleteUser          = "Cannot delete user. "

	ErrorBindInput      = "Cannot bind input data. "
	ErrorNamepassEmpty  = "\"username\" and \"password\" are required fields, and one or both of them cannot be empty."
	ErrorCheckExistence = "Cannot check user existence. "
	ErrorIdConverting   = "Cannot convert requested param \"ID\" to type \"int\". "
)
