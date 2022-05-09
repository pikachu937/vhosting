package messages

import (
	"github.com/mikerumy/vhosting/internal/models"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func InfoUserCreated() *models.Log {
	return &models.Log{StatusCode: 200, Message: "User created.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotUserData(usr *models.User) *models.Log {
	return &models.Log{StatusCode: 200, Message: usr, ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotAllUsersData(users map[int]*models.User) *models.Log {
	return &models.Log{StatusCode: 200, Message: users, ErrorLevel: logger.ErrLevelInfo}
}

func InfoUserPartiallyUpdated() *models.Log {
	return &models.Log{StatusCode: 200, Message: "User partially updated.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoUserDeleted() *models.Log {
	return &models.Log{StatusCode: 200, Message: "User deleted.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoYouHaveSuccessfullySignedIn() *models.Log {
	return &models.Log{StatusCode: 202, Message: "You have successfully signed-in.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoYouHaveSuccessfullyChangedPassword() *models.Log {
	return &models.Log{StatusCode: 202, Message: "You have successfully changed password.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoYouHaveSuccessfullySignedOut() *models.Log {
	return &models.Log{StatusCode: 202, Message: "You have successfully signed-out.", ErrorLevel: logger.ErrLevelInfo}
}
