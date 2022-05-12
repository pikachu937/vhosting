package usergroup

type UGRepository interface {
	UGCommon

	CreateUsergroup(userId, groupId int) error
	UpdateUsergroup(userId, groupId int) error
}
