package messages

import (
	"github.com/deepch/vdk/av"
	lg "github.com/mikerumy/vhosting/internal/logging"
	"github.com/mikerumy/vhosting/pkg/logger"
)

func ErrorTrackIsIgnoredCodecNotSupportedWebRTC(codecType av.CodecType) *lg.Log {
	return &lg.Log{ErrCode: 900, Message: "Track is ignored - codec not supported WebRTC. codec type: " + codecType.String(), ErrLevel: logger.ErrLevelError}
}

func ErrorWritingOfCodecError(err error) *lg.Log {
	return &lg.Log{ErrCode: 901, Message: "Writing of codec error. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoStreamNotFound(suuid string) *lg.Log {
	return &lg.Log{Message: "Stream not found. Suuid: " + suuid}
}

func InfoStreamCodecNotFound(suuid string) *lg.Log {
	return &lg.Log{Message: "Stream codec not found. Suuid: " + suuid}
}

func ErrorWriteHeaderError(err error) *lg.Log {
	return &lg.Log{ErrCode: 902, Message: "WriteHeader error. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorCannotWriteBytes(err error) *lg.Log {
	return &lg.Log{ErrCode: 903, Message: "Cannot write bytes. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorStreamCodecNotFound(err error) *lg.Log {
	return &lg.Log{ErrCode: 904, Message: "Stream codec not found. LastError: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorMuxerWriteHeaderError(err error) *lg.Log {
	return &lg.Log{ErrCode: 905, Message: "Muxer WriteHeader error. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoStreamTriesToConnect(name string) *lg.Log {
	return &lg.Log{Message: "Stream tries to connect " + name}
}

func ErrorRTSPWorkerError(err error) *lg.Log {
	return &lg.Log{ErrCode: 906, Message: "rtspWorker error. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorOnDemandANDNotHasViewerError(errmsg string) *lg.Log {
	return &lg.Log{ErrCode: 907, Message: "onDemand && notHasViewer error. Error: " + errmsg, ErrLevel: logger.ErrLevelError}
}

func ErrorFrameDecoderSingleError(err error) *lg.Log {
	return &lg.Log{ErrCode: 908, Message: "frameDecoderSingle error. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func InfoSnapshotCreated(name string) *lg.Log {
	return &lg.Log{Message: "Snapshot for " + name + " created"}
}

func ErrorBadVideoCodecWaitingForSPS_PPS() *lg.Log {
	return &lg.Log{ErrCode: 909, Message: "Bad video codec - waiting for SPS/PPS", ErrLevel: logger.ErrLevelError}
}

func InfoNoVideo() *lg.Log {
	return &lg.Log{Message: "No video"}
}

func ErrorWritePacketError(err error) *lg.Log {
	return &lg.Log{ErrCode: 910, Message: "WritePacket error. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}

func ErrorPseudoUUIDReadError(err error) *lg.Log {
	return &lg.Log{ErrCode: 911, Message: "pseudoUUID read error. Error: " + err.Error(), ErrLevel: logger.ErrLevelError}
}
