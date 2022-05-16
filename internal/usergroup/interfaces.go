package usergroup

type UGCommon interface {
	CreateUsergroup(userId, groupId int) error
	DeleteUsergroup(userId, groupId int) error
}

type UGUseCase interface {
	UGCommon
}

type UGRepository interface {
	UGCommon
}
