package cpu_temp

import (
	"errors"
	"github.com/ssimunic/gosensors"
)

type Stats struct {
	Temperature string
}

func Get() (*Stats, error) {
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		return nil, err
	}

	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
			if key == "CPU" {
				return &Stats{Temperature: value}, nil
			}
		}
	}

	return nil, errors.New("not found chip")
}
