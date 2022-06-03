package messages

import (
	"fmt"

	"github.com/deepch/vdk/av"
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorStreamCodecNotFound(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 900, Message: fmt.Sprintf("Stream codec not found. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorMuxerWriteHeaderError(err error) *lg.Log {
	return &lg.Log{StatusCode: 500, ErrCode: 901, Message: fmt.Sprintf("Muxer WriteHeader error. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorTrackIsIgnoredBecauseCodecNotSupportedWebRTC(codecType av.CodecType) *lg.Log {
	return &lg.Log{ErrCode: 902, Message: fmt.Sprintf("Track is ignored - codec not supported WebRTC. CodecType: %v.", codecType), ErrLevel: logger.ErrLevelError}
}

func ErrorWritingOfCodecInfoError(err error) *lg.Log {
	return &lg.Log{ErrCode: 903, Message: fmt.Sprintf("Writing of codec Info error. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoStreamNotFound(suuid string) *lg.Log {
	return &lg.Log{Message: fmt.Sprintf("Stream not found. Stream UUID: %s.", suuid), ErrLevel: logger.ErrLevelInfo}
}

func InfoStreamCodecNotFound(suuid string) *lg.Log {
	return &lg.Log{Message: "Stream codec not found.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorWriteHeaderError(err error) *lg.Log {
	return &lg.Log{ErrCode: 904, Message: fmt.Sprintf("WriteHeader error. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotWriteBytesError(err error) *lg.Log {
	return &lg.Log{ErrCode: 905, Message: fmt.Sprintf("Cannot write bytes error. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func InfoNoVideo() *lg.Log {
	return &lg.Log{Message: "No video.", ErrLevel: logger.ErrLevelInfo}
}

func ErrorWritePacketError(err error) *lg.Log {
	return &lg.Log{ErrCode: 906, Message: fmt.Sprintf("Write packet error. Error: %s.", err.Error()), ErrLevel: logger.ErrLevelError}
}

func ErrorBadVideoCodecWaitingForSPS_PPS() *lg.Log {
	return &lg.Log{ErrCode: 907, Message: "Bad video codec - waiting for SPS/PPS", ErrLevel: logger.ErrLevelError}
}

func InfoCalledFunction(signature string) *lg.Log {
	return &lg.Log{ErrCode: 908, Message: "Bad video codec - waiting for SPS/PPS", ErrLevel: logger.ErrLevelError}
}
