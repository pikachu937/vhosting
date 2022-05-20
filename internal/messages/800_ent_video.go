package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/video"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorUrlAndFilenameCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 800, Message: fmt.Sprintf("URL and file name cannot be empty."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 801, Message: fmt.Sprintf("Cannot create video. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoVideoCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video created.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotCheckVideoExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 802, Message: fmt.Sprintf("Cannot check video existence. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorVideoWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 803, Message: fmt.Sprintf("Video with requested ID is not exist."), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 804, Message: fmt.Sprintf("Cannot get video. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoGotVideo(nfo *video.Video) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: nfo, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetAllVideos(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 805, Message: fmt.Sprintf("Cannot get all videos. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoNoVideosAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No videos available.", ErrLevel: logger.ErrLevelInfo}
}

func InfoGotAllVideos(users map[int]*video.Video) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: users, ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotPartiallyUpdateVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 806, Message: fmt.Sprintf("Cannot partially update video. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoVideoPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video partially updated.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 807, Message: fmt.Sprintf("Cannot delete video. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoVideoDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video deleted.", ErrLevel: logger.ErrLevelInfo}
}
