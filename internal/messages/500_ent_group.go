package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/group"
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorGroupNameCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 500, Message: fmt.Sprintf("Group name cannot be empty."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCheckGroupExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 501, Message: fmt.Sprintf("Cannot check group existence. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorGroupWithEnteredNameIsExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 502, Message: "Group with entered name is exist.", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 503, Message: fmt.Sprintf("Cannot create group. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGroupCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group created.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorGroupWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 504, Message: fmt.Sprintf("Group with requested ID is not exist."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 505, Message: fmt.Sprintf("Cannot get group. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGotGroup(grp *group.Group) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: grp, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetAllGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 506, Message: fmt.Sprintf("Cannot get all groups. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoNoGroupsAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No groups available.", ErrLevel: logger.ErrLevelInfo}
}

func InfoGotAllGroups(groups map[int]*group.Group) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groups, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotPartiallyUpdateGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 507, Message: fmt.Sprintf("Cannot partially update group. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGroupPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group partially updated.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 508, Message: fmt.Sprintf("Cannot delete group. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGroupDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group deleted.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorGroupIdsCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 509, Message: fmt.Sprintf("Group IDs cannot be empty."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotSetUserGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 510, Message: fmt.Sprintf("Cannot set user groups. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoUserGroupsSet() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User groups set.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetUserGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 511, Message: fmt.Sprintf("Cannot get user groups. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGotUserGroups(groupIds *group.GroupIds) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groupIds, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteUserGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 512, Message: fmt.Sprintf("Cannot delete user groups. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoUserGroupsDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "User groups deleted.", ErrLevel: logger.ErrLevelInfo}
}
