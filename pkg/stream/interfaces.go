package stream

import (
	"github.com/deepch/vdk/av"
	webrtc "github.com/deepch/vdk/format/webrtcv3"
)

type StreamUseCase interface {
	ServeStreams()
	Exit(suuid string) bool
	RunIfNotRun(uuid string)
	CodecGet(suuid string) []av.CodecData
	GetICEServers() []string
	GetICEUsername() string
	GetICECredential() string
	GetWebRTCPortMin() uint16
	GetWebRTCPortMax() uint16
	WritePackets(url string, muxerWebRTC *webrtc.Muxer, audioOnly bool)
	CastListAdd(suuid string) (string, chan av.Packet)
	CastListDelete(suuid, cuuid string)
	List() (string, []string)
}
