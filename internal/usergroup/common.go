package usergroup

type UGCommon interface {
	IsUserInGroup(userId, groupId int) (bool, error)
}
