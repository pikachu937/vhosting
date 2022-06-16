package config_stream

import (
	"sync"

	"github.com/mikerumy/vhosting/pkg/stream"
)

type Config struct {
	Mutex     sync.RWMutex
	Server    stream.Server                    `json:"server"`
	Streams   map[string]stream.StreamSettings `json:"streams"`
	LastError error
}
