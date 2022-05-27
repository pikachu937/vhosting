package models

import (
	"sync"

	"github.com/deepch/vdk/av"
)

type ConfigST struct {
	Mutex     sync.RWMutex
	Server    Server            `json:"server"`
	Streams   map[string]Stream `json:"streams"`
	LastError error
}

type Server struct {
	HTTPPort      string
	ICEServers    []string `json:"ice_servers"`
	ICEUsername   string   `json:"ice_username"`
	ICECredential string   `json:"ice_credential"`
	WebRTCPortMin uint16   `json:"webrtc_port_min"`
	WebRTCPortMax uint16   `json:"webrtc_port_max"`
}

type Stream struct {
	URL          string `json:"url"`
	Status       bool   `json:"status"`
	OnDemand     bool   `json:"on_demand"`
	DisableAudio bool   `json:"disable_audio"`
	Debug        bool   `json:"debug"`
	RunLock      bool   `json:"-"`
	Codecs       []av.CodecData
	ClientList   map[string]Viewer
}

type Viewer struct {
	Cast chan av.Packet
}
