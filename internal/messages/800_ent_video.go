package messages

import (
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/internal/video"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorUrlAndFilenameCannotBeEmpty() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 800, Message: "URL and file name cannot be empty", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotCreateVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 801, Message: "Cannot create video. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoVideoCreated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video created"}
}

func ErrorCannotCheckVideoExistence(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 802, Message: "Cannot check video existence. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorVideoWithRequestedIDIsNotExist() *lg.Log {
	return &lg.Log{StatusCode: 400, ErrCode: 803, Message: "Video with requested ID is not exist", ErrLevel: logger.ErrLevelError}
}

func ErrorCannotGetVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 804, Message: "Cannot get video. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoGotVideo(nfo *video.Video) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: nfo}
}

func ErrorCannotGetAllVideos(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 805, Message: "Cannot get all videos. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoNoVideosAvailable() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "No videos available"}
}

func InfoGotAllVideos(users map[int]*video.Video) *lg.Log {
	return &lg.Log{StatusCode: 200, Message: users}
}

func ErrorCannotPartiallyUpdateVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 806, Message: "Cannot partially update video. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoVideoPartiallyUpdated() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video partially updated"}
}

func ErrorCannotDeleteVideo(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 807, Message: "Cannot delete video. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoVideoDeleted() *lg.Log {
	return &lg.Log{StatusCode: 200, Message: "Video deleted"}
}
