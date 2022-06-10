package group

type Group struct {
	Id   int    `json:"id"   db:"id"`
	Name string `json:"name" db:"name"`
}

type GroupIds struct {
	Ids []int `json:"groupIds" db:"group_id"`
}
