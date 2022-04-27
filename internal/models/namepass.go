package models

type Namepass struct {
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password" db:"password_hash"`
}
