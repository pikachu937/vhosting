package vhs

type User struct {
	Id    int    `json:"-" db:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
