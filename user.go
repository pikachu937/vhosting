package vh

type User struct {
	Id           int    `json:"id"           db:"id"`
	Username     string `json:"username"     db:"username"`
	PasswordHash string `json:"password"     db:"password_hash"`
	IsActive     bool   `json:"is-active"    db:"is_active"`
	IsSuperUser  bool   `json:"is-superuser" db:"is_superuser"`
	IsStaff      bool   `json:"is-staff"     db:"is_staff"`
	FirstName    string `json:"first-name"   db:"first_name"`
	LastName     string `json:"last-name"    db:"last_name"`
	DateJoined   string `json:"date-joined"  db:"date_joined"`
	LastLogin    string `json:"last-login"   db:"last_login"`
}
