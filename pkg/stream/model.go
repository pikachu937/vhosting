package stream

import (
	"database/sql"
)

type Stream struct {
	Id           int    `json:"id" db:"id"`
	Stream       string `json:"stream" db:"Stream"`
	DateTime     string `json:"dateTime" db:"DateTime"`
	StatePublic  int    `json:"-" db:"StatePublic"`
	StatusPublic int    `json:"statusPublic" db:"StatusPublic"`
	StatusRecord int    `json:"-" db:"StatusRecord"`
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

type JCodec struct {
	Type string
}

type Response struct {
	Tracks []string `json:"tracks"`
	Sdp64  string   `json:"sdp64"`
}
