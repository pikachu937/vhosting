package auth

type Namepass struct {
	Id           int    `json:"id"       db:"id"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password" db:"password_hash"`
}
