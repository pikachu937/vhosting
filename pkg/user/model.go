package user

type User struct {
	Id           int    `json:"id"           db:"id"`
	Username     string `json:"username"     db:"username"`
	PasswordHash string `json:"password"     db:"password_hash"`
	IsActive     bool   `json:"isActive"    db:"is_active"`
	IsSuperuser  bool   `json:"isSuperuser" db:"is_superuser"`
	IsStaff      bool   `json:"isStaff"     db:"is_staff"`
	FirstName    string `json:"firstName"   db:"first_name"`
	LastName     string `json:"lastName"    db:"last_name"`
	JoiningDate  string `json:"joiningDate" db:"joining_date"`
	LastLogin    string `json:"lastLogin"   db:"last_login"`
}

type Pagin struct {
	Limit int
	Page  int
}
