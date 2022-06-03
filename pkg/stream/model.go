package stream

import "github.com/deepch/vdk/av"

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

type JCodec struct {
	Type string
}

type Response struct {
	Tracks []string `json:"tracks"`
	Sdp64  string   `json:"sdp64"`
}
