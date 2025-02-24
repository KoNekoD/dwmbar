package cpu_temp

import (
	"errors"
	"fmt"
	"github.com/ssimunic/gosensors"
	"regexp"
)

type Stats struct {
	Temperature string
}

func Get() (*Stats, error) {
	sensors, err := gosensors.NewFromSystem()
	if err != nil {
		return nil, fmt.Errorf("error getting sensors data: %v", err)
	}

	packageTempRegex := regexp.MustCompile(`([+-]?\d+(\.\d+)?)°?C`)
	for _, values := range sensors.Chips {
		for key, value := range values {
			if key == "CPU" {
				return &Stats{Temperature: value}, nil
			} else if key == "Package id 0" {
				matches := packageTempRegex.FindStringSubmatch(value)
				if len(matches) > 0 {
					return &Stats{Temperature: matches[0]}, nil
				}
			}
		}
	}

	return nil, errors.New("not found chip or temperature data")
}
