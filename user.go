package vhs

type User struct {
	Id           int    `json:"-" db:"id"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password" db:"password_hash"`
}
