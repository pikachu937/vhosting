package dbconsts

const (
	TableUsers  = "users"
	Id          = "id"
	Username    = "username"
	PassHash    = "password_hash"
	IsActive    = "is_active"
	IsSuperUser = "is_superuser"
	IsStaff     = "is_staff"
	FirstName   = "first_name"
	LastName    = "last_name"
	JoiningDate = "joining_date"
	LastLogin   = "last_login"

	TableSessions = "sessions"
	Content       = "content"
	CreationDate  = "creation_date"

	TableLogs     = "logs"
	SessionOwner  = "session_owner"
	RequestMethod = "request_method"
	RequestPath   = "request_path"
	StatusCode    = "status_code"
	ErrorCode     = "error_code"
	Message       = "message"

	TableUserGroups = "user_groups"
	UserId          = "user_id"
	GroupId         = "group_id"
)
