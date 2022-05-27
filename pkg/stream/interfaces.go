package stream

import (
	"github.com/deepch/vdk/av"
)

type StreamUseCase interface {
	ServeStreams()
	RunIfNotRun(uuid string)
	GetICEServers() []string
	GetICEUsername() string
	GetICECredential() string
	GetWebRTCPortMin() uint16
	GetWebRTCPortMax() uint16
	List() (string, []string)
	Exit(suuid string) bool
	CodecGet(suuid string) []av.CodecData
	CastListAdd(suuid string) (string, chan av.Packet)
	CastListDelete(suuid, cuuid string)
}
