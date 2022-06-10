package permission

type Perm struct {
	Id       int    `json:"id"        db:"id"`
	Name     string `json:"name"      db:"name"`
	Codename string `json:"codeName" db:"code_name"`
}

type PermIds struct {
	Ids []int `json:"permIds" db:"perm_id"`
}
