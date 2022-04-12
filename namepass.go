package vh

type NamePass struct {
	Username     string `json:"username" binding:"required" db:"username"`
	PasswordHash string `json:"password" binding:"required" db:"password_hash"`
}
