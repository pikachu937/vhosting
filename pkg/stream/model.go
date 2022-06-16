package stream

import (
	"database/sql"

	"github.com/deepch/vdk/av"
)

type Stream struct {
	Id           int    `json:"id" db:"id"`
	Stream       string `json:"stream" db:"Stream"`
	DateTime     string `json:"dateTime" db:"DateTime"`
	StatePublic  int    `json:"statePublic" db:"StatePublic"`
	StatusPublic int    `json:"statusPublic" db:"StatusPublic"`
	StatusRecord int    `json:"statusRecord" db:"StatusRecord"`
	PathStream   string `json:"pathStream" db:"pathStream"`
}

type StreamGet struct {
	Id           int            `db:"id"`
	Stream       sql.NullString `db:"Stream"`
	DateTime     sql.NullString `db:"DateTime"`
	StatePublic  sql.NullInt16  `db:"StatePublic"`
	StatusPublic sql.NullInt16  `db:"StatusPublic"`
	StatusRecord sql.NullInt16  `db:"StatusRecord"`
	PathStream   sql.NullString `db:"pathStream"`
}

type Server struct {
	HTTPPort      string
	ICEServers    []string `json:"iceServers"`
	ICEUsername   string   `json:"iceUsername"`
	ICECredential string   `json:"iceCredential"`
	WebRTCPortMin uint16   `json:"webrtcPortMin"`
	WebRTCPortMax uint16   `json:"webrtcPortMax"`
}

type StreamSettings struct {
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
