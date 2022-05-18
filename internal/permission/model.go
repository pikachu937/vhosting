package permission

type PermIds struct {
	Ids []int `json:"perm-ids" db:"perm_id"`
}

type Perm struct {
	Id       int    `json:"id"        db:"id"`
	Name     string `json:"name"      db:"name"`
	Codename string `json:"code-name" db:"code_name"`
}
