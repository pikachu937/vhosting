package user

import "github.com/gin-gonic/gin"

type UserCommon interface {
	CreateUser(usr User) error
	GetUser(id int) (*User, error)
	GetAllUsers() (map[int]*User, error)
	PartiallyUpdateUser(usr *User) error
	DeleteUser(id int) error

	IsUserSuperuserOrStaff(username string) (bool, error)
	IsUserHavePersonalPermission(userId int, userPerm string) (bool, error)
	IsUserExists(idOrUsername interface{}) (bool, error)
	GetUserId(username string) (int, error)
}

type UserUseCase interface {
	UserCommon

	IsRequiredEmpty(username, password string) bool
	BindJSONUser(ctx *gin.Context) (User, error)
	AtoiRequestedId(ctx *gin.Context) (int, error)
}

type UserRepository interface {
	UserCommon
}
