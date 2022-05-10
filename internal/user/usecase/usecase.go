package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/user"
	"github.com/mikerumy/vhosting/pkg/config_tool"
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

func (u *UserUseCase) CreateUser(ctx *gin.Context, usr user.User, timestamp string) error {
	usr.JoiningDate = timestamp
	usr.LastLogin = timestamp
	usr.IsActive = true
	return u.userRepo.CreateUser(usr)
}

func (u *UserUseCase) GetUser(id int) (*user.User, error) {
	return u.userRepo.GetUser(id)
}

func (u *UserUseCase) GetAllUsers() (map[int]*user.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserUseCase) PartiallyUpdateUser(id int, usr user.User) error {
	return u.userRepo.PartiallyUpdateUser(id, usr)
}

func (u *UserUseCase) DeleteUser(id int) error {
	return u.userRepo.DeleteUser(id)
}

func (u *UserUseCase) IsUserExists(idOrUsername interface{}) (bool, error) {
	exists, err := u.userRepo.IsUserExists(idOrUsername)
	if err != nil {
		return false, err
	}
	return exists, nil
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
