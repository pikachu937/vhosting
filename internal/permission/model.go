package permission

type Permission struct {
	Id       int    `json:"id"        db:"id"`
	Name     string `json:"name"      db:"name"`
	CodeName string `json:"code-name" db:"code_name"`
}
