package stream

import "github.com/deepch/vdk/av"

type Server struct {
	HTTPPort      string
	ICEServers    []string `json:"iceServers"`
	ICEUsername   string   `json:"iceUsername"`
	ICECredential string   `json:"iceCredential"`
	WebRTCPortMin uint16   `json:"webrtcPortMin"`
	WebRTCPortMax uint16   `json:"webrtcPortMax"`
}

type Stream struct {
	URL          string `json:"url"`
	Status       bool   `json:"status"`
	OnDemand     bool   `json:"onDemand"`
	DisableAudio bool   `json:"disableAudio"`
	Debug        bool   `json:"debug"`
	RunLock      bool   `json:"-"`
	Codecs       []av.CodecData
	ClientList   map[string]Viewer
}

type Viewer struct {
	Cast chan av.Packet
}

type JCodec struct {
	Type string
}

type Response struct {
	Tracks []string `json:"tracks"`
	Sdp64  string   `json:"sdp64"`
}
