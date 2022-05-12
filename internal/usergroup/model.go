package usergroup

type Usergroup struct {
	Id      int    `db:"id"`
	UserId  string `db:"user_id"`
	GroupId string `db:"group_id"`
}
