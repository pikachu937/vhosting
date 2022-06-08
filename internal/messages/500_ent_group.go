package messages

import (
	"github.com/mikerumy/vhosting/internal/group"
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorGroupNameCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 500, Message: "Group name cannot be empty", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckGroupExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 501, Message: "Cannot check group existence. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorGroupWithEnteredNameIsExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 502, Message: "Group with entered name is exist", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 503, Message: "Cannot create group. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGroupCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group created"}
}

func ErrorGroupWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 504, Message: "Group with requested ID is not exist", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 505, Message: "Cannot get group. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGotGroup(grp *group.Group) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: grp}
}

func ErrorCannotGetAllGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 506, Message: "Cannot get all groups. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoNoGroupsAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No groups available"}
}

func InfoGotAllGroups(groups map[int]*group.Group) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groups}
}

func ErrorCannotPartiallyUpdateGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 507, Message: "Cannot partially update group. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGroupPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group partially updated"}
}

func ErrorCannotDeleteGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 508, Message: "Cannot delete group. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGroupDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group deleted"}
}

func ErrorGroupIdsCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 509, Message: "Group IDs cannot be empty", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotSetUserGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 510, Message: "Cannot set user groups. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoUserGroupsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User groups set"}
}

func ErrorCannotGetUserGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 511, Message: "Cannot get user groups. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGotUserGroups(groupIds *group.GroupIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groupIds}
}

func ErrorCannotDeleteUserGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 512, Message: "Cannot delete user groups. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoUserGroupsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User groups deleted"}
}
