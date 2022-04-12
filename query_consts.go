package vh

const (
	UsersTable  = "users"
	Id          = "id"
	Username    = "username"
	PassHash    = "password_hash"
	IsActive    = "is_active"
	IsSuperUser = "is_superuser"
	IsStaff     = "is_staff"
	FirstName   = "first_name"
	LastName    = "last_name"
	Email       = "email"
	DateJoined  = "date_joined"
	LastLogin   = "last_login"

	SessionsTable = "sessions"
	// Id = "id"
	Content      = "content"
	CreationDate = "creation_date"

	INSERT_INTO_TBL_VALUES_VAL_RETURNING_RET = "INSERT INTO %s VALUES %s RETURNING %s"
	INSERT_INTO_TBL_VALUES_VAL               = "INSERT INTO %s VALUES %s"
	SELECT_COL_FROM_TBL_WHERE_CND            = "SELECT %s FROM %s WHERE %s"
	SELECT_COL_FROM_TBL                      = "SELECT %s FROM %s"
	UPDATE_TBL_SET_VAL_WHERE_CND             = "UPDATE %s SET %s WHERE %s"
	DELETE_FROM_TBL_WHERE_CND                = "DELETE FROM %s WHERE %s"
	DELETE_CASCADE_FROM_TBL_WHERE_CND        = "DELETE CASCADE FROM %s WHERE %s"
)
