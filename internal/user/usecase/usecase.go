package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/config_tool"
	"github.com/mikerumy/vhosting/pkg/cookie_tool"
	"github.com/mikerumy/vhosting/pkg/hasher"
)

type UserUseCase struct {
	cfg      config_tool.Config
	userRepo user.UserRepository
}

func NewUserUseCase(cfg config_tool.Config, userRepo user.UserRepository) *UserUseCase {
	return &UserUseCase{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (u *UserUseCase) CreateUser(usr user.User) error {
	usr.IsActive = true
	usr.IsSuperuser = false
	usr.IsStaff = false
	usr.LastLogin = usr.JoiningDate
	return u.userRepo.CreateUser(usr)
}

func (u *UserUseCase) GetUser(id int) (*user.User, error) {
	return u.userRepo.GetUser(id)
}

func (u *UserUseCase) GetAllUsers() (map[int]*user.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserUseCase) PartiallyUpdateUser(usr *user.User) error {
	return u.userRepo.PartiallyUpdateUser(usr)
}

func (u *UserUseCase) DeleteUser(id int) error {
	return u.userRepo.DeleteUser(id)
}

func (u *UserUseCase) IsRequiredEmpty(username, password string) bool {
	if username == "" || password == "" {
		return true
	}
	return false
}

func (u *UserUseCase) IsUserSuperuserOrStaff(username string) (bool, error) {
	return u.userRepo.IsUserSuperuserOrStaff(username)
}

func (u *UserUseCase) IsUserHavePersonalPermission(userId int, userPerm string) (bool, error) {
	return u.userRepo.IsUserHavePersonalPermission(userId, userPerm)
}

func (u *UserUseCase) IsUserExists(idOrUsername interface{}) (bool, error) {
	exists, err := u.userRepo.IsUserExists(idOrUsername)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserUseCase) GetUserId(username string) (int, error) {
	return u.userRepo.GetUserId(username)
}

func (u *UserUseCase) BindJSONUser(ctx *gin.Context) (user.User, error) {
	var usr user.User
	if err := ctx.BindJSON(&usr); err != nil {
		return usr, err
	}
	if usr.PasswordHash != "" {
		usr.PasswordHash = hasher.GeneratePasswordHash(usr.PasswordHash, u.cfg.HashingPasswordSalt)
	}
	return usr, nil
}

func (u *UserUseCase) AtoiRequestedId(ctx *gin.Context) (int, error) {
	idInt, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return -1, err
	}
	return idInt, nil
}

func (u *UserUseCase) ReadCookie(ctx *gin.Context) string {
	return cookie_tool.Read(ctx)
}
