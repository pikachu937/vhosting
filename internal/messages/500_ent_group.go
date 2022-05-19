package messages

import (
	"fmt"

	"github.com/mikerumy/vhosting/internal/group"
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorGroupNameCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 500, Message: fmt.Sprintf("Group name cannot be empty."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCheckGroupExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 501, Message: fmt.Sprintf("Cannot check group existence. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorGroupWithEnteredNameIsExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 502, Message: "Group with entered name is exist.", ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCreateGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 503, Message: fmt.Sprintf("Cannot create group. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGroupCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group created.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorGroupWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 504, Message: fmt.Sprintf("Group with requested ID is not exist."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGetGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 505, Message: fmt.Sprintf("Cannot get group. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGotGroupData(grp *group.Group) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: grp, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetAllGroups(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 506, Message: fmt.Sprintf("Cannot get all groups. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoNoGroupsAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No groups available.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotAllGroupsData(groups map[int]*group.Group) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: groups, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotPartiallyUpdateGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 507, Message: fmt.Sprintf("Cannot partially update group. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGroupPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group partially updated.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteGroup(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 508, Message: fmt.Sprintf("Cannot delete group. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGroupDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Group deleted.", ErrorLevel: logger.ErrLevelInfo}
}
