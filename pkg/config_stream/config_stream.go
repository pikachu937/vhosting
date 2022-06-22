package config_stream

func LoadConfig(path string) (*SConfig, error) {
	var scfg SConfig

	url := "43704893903143017940"
	scfg.Streams = map[string]Stream{}
	scfg.Streams[url] = Stream{ClientList: make(map[string]Viewer), URL: url}

	return &scfg, nil
}
