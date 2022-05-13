package userperm

type Userperm struct {
	Id     int    `db:"id"`
	UserId string `db:"user_id"`
	PermId string `db:"perm_id"`
}
