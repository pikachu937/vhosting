package session

type Session struct {
	Id           int    `db:"id"`
	Content      string `db:"content"`
	CreationDate string `db:"creationDate"`
}
