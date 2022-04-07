package vh

type User struct {
	Id           int    `json:"-" db:"id"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password-hash" db:"password_hash"`
}
