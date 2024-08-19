package volume_state

import (
	"github.com/itchyny/volume-go"
)

type Stats struct {
	Volume int
	Muted  bool
}

func Get() (*Stats, error) {
	muted, err := volume.GetMuted()

	if err != nil {
		return nil, err
	}

	vol, err2 := volume.GetVolume()
	if err2 != nil {
		return nil, err2
	}

	return &Stats{
		Volume: vol,
		Muted:  muted,
	}, nil
}
