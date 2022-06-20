package config_stream

func LoadConfig(path string) (*SConfig, error) {
	var scfg SConfig

	// data, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	udpMin := flag.Int("udp_min", 0, "WebRTC UDP port min")
	// 	udpMax := flag.Int("udp_max", 0, "WebRTC UDP port max")
	// 	flag.Parse()

	// 	scfg.Server.WebRTCPortMin = uint16(*udpMin)
	// 	scfg.Server.WebRTCPortMax = uint16(*udpMax)

	// 	scfg.Streams = make(map[string]Stream)
	// }

	// err = json.Unmarshal(data, &scfg)
	// if err != nil {
	// 	return nil, err
	// }

	// for i, val := range scfg.Streams {
	// 	val.ClientList = make(map[string]Viewer)
	// 	i = val.URL
	// 	scfg.Streams[i] = val
	// }

	urls := []string{"43704893903143017940", "21854368092658842199",
		"41574336307958053922", "93884975412010002133"}

	scfg.Streams = map[string]Stream{}
	for _, val := range urls {
		clientlist := make(map[string]Viewer)
		scfg.Streams[val] = Stream{ClientList: clientlist, URL: val}
	}

	return &scfg, nil
}
