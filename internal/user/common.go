package user

type UserCommon interface {
	GetUser(id int) (*User, error)
	GetAllUsers() (map[int]*User, error)
	PartiallyUpdateUser(id int, usr User) error
	DeleteUser(id int) error

	IsUserExists(idOrUsername interface{}) (bool, error)
}
