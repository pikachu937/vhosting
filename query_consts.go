package vh

const (
	Id                            = "id"
	UsersTable                    = "users"
	Username                      = "username"
	PassHash                      = "password_hash"
	INSERT_TBL_COL_VAL_RET        = "INSERT INTO %s %s VALUES %s RETURNING %s"
	SELECT_COL_FROM_TBL_WHERE_CND = "SELECT %s FROM %s WHERE %s"
	SELECT_COL_FROM_TBL           = "SELECT %s FROM %s"
	UPDATE_TBL_SET_VAL_WHERE_CND  = "UPDATE users SET username=$1, password_hash=$2 WHERE id=$3"
)
