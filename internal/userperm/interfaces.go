package userperm

type UPCommon interface {
	CreateUserperm(userperm *Userperm) error
	GetUserPermissions(id int) (map[int]*Userperm, error)
	UpsertUserPermissions(userId, groupId int) error
	DeleteUserPermissions(userId, groupId int) error
}

type UPUseCase interface {
	UPCommon
}

type UPRepository interface {
	UPCommon
}
