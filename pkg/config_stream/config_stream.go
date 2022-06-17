package config_stream

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	"github.com/mikerumy/vhosting/pkg/config"
)

func LoadConfig(path string) (*SConfig, error) {
	var scfg SConfig

	data, err := ioutil.ReadFile(path)
	if err != nil {
		udpMin := flag.Int("udp_min", 0, "WebRTC UDP port min")
		udpMax := flag.Int("udp_max", 0, "WebRTC UDP port max")
		flag.Parse()

		scfg.Server.WebRTCPortMin = uint16(*udpMin)
		scfg.Server.WebRTCPortMax = uint16(*udpMax)

		scfg.Streams = make(map[string]Stream)
	}

	err = json.Unmarshal(data, &scfg)
	if err != nil {
		return nil, err
	}

	for i, val := range scfg.Streams {
		val.ClientList = make(map[string]Viewer)
		scfg.Streams[i] = val
	}

	return &scfg, nil
}

func LoadRTSPLink(cfg *config.Config) (*SConfig, error) {
	var scfg SConfig
	var strm Stream

	strm.DisableAudio = true
	strm.URL = cfg.StreamLink

	for i, val := range scfg.Streams {
		val.ClientList = make(map[string]Viewer)
		scfg.Streams[i] = val
	}

	return &scfg, nil
}
