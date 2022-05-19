package messages

import (
	"fmt"

	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/video"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorUrlAndFilenameCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 800, Message: fmt.Sprintf("URL and file name cannot be empty."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotCreateVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 801, Message: fmt.Sprintf("Cannot create video. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoVideoCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video created.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotCheckVideoExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 802, Message: fmt.Sprintf("Cannot check video existence. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func ErrorVideoWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrorCode: 803, Message: fmt.Sprintf("Video with requested ID is not exist."), ErrorLevel: logger.ErrLevelError}
}

func ErrorCannotGetVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 804, Message: fmt.Sprintf("Cannot get video. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoGotVideo(nfo *video.Video) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: nfo, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotGetAllVideos(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 805, Message: fmt.Sprintf("Cannot get all videos. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoNoVideosAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No videos available.", ErrorLevel: logger.ErrLevelInfo}
}

func InfoGotAllVideos(users map[int]*video.Video) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: users, ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotPartiallyUpdateVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 806, Message: fmt.Sprintf("Cannot partially update video. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoVideoPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video partially updated.", ErrorLevel: logger.ErrLevelInfo}
}

func ErrorCannotDeleteVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrorCode: 807, Message: fmt.Sprintf("Cannot delete video. Error: %s.", err.Error()), ErrorLevel: logger.ErrLevelError}
}

func InfoVideoDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video deleted.", ErrorLevel: logger.ErrLevelInfo}
}
