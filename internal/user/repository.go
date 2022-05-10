package user

type UserRepository interface {
	UserCommon

	CreateUser(usr User) error
}
