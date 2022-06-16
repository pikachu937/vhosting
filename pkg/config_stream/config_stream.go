package config_stream

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/mikerumy/vhosting/pkg/stream"
)

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		addr := flag.String("listen", "8000", "HTTP host:port")
		udpMin := flag.Int("udp_min", 0, "WebRTC UDP port min")
		udpMax := flag.Int("udp_max", 0, "WebRTC UDP port max")
		iceServer := flag.String("ice_server", "", "ICE Server")
		flag.Parse()

		cfg.Server.HTTPPort = *addr
		cfg.Server.WebRTCPortMin = uint16(*udpMin)
		cfg.Server.WebRTCPortMax = uint16(*udpMax)
		if len(*iceServer) > 0 {
			cfg.Server.ICEServers = []string{*iceServer}
		}

		cfg.Streams = make(map[string]stream.StreamSettings)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	for i, val := range cfg.Streams {
		val.ClientList = make(map[string]stream.Viewer)
		cfg.Streams[i] = val
	}

	return &cfg, nil
}
