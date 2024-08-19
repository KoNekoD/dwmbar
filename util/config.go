package util

import "flag"

type Config struct {
	NoNetworkStats   bool
	NoCpu            bool
	NoTemp           bool
	NoMemory         bool
	NoKeyboardLayout bool
	NoVolume         bool
	NoBrightness     bool
	NoNetworkState   bool
	NoPowerState     bool
}

func NewConfig() *Config {
	noNetworkStats := flag.Bool("noNetworkStats", false, "no network stat")
	noCpu := flag.Bool("noCpu", false, "no cpu stat")
	noTemp := flag.Bool("noTemp", false, "no cpu temp")
	noMemory := flag.Bool("noMemory", false, "no memory")
	noLang := flag.Bool("noLang", false, "no lang")
	noVolume := flag.Bool("noVolume", false, "no volume")
	noBrightness := flag.Bool("noBrightness", false, "no brightness")
	noNetworkState := flag.Bool("noNetworkState", false, "no network state")
	noPowerState := flag.Bool("noPowerState", false, "no power state")

	flag.Parse()

	return &Config{
		NoNetworkStats:   *noNetworkStats,
		NoCpu:            *noCpu,
		NoTemp:           *noTemp,
		NoMemory:         *noMemory,
		NoKeyboardLayout: *noLang,
		NoVolume:         *noVolume,
		NoBrightness:     *noBrightness,
		NoNetworkState:   *noNetworkState,
		NoPowerState:     *noPowerState,
	}
}
