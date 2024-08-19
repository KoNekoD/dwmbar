package brightness_state

import (
	"errors"
	"main/util"
	"regexp"
	"strconv"
	"strings"
)

type Stats struct {
	Brightness int
}

func Get() (*Stats, error) {
	brightness, err := getBrightnessState()
	if err != nil {
		return nil, err
	}

	return &Stats{
		Brightness: brightness,
	}, nil
}

func getBrightnessState() (int, error) {
	output, err := util.ExecCmd("brightnessctl")
	if err != nil {
		return 0, err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		var regex = regexp.MustCompile("([^()]+)%")

		found := regex.FindStringSubmatch(line)
		if len(found) > 0 {
			code, err2 := strconv.Atoi(found[1])

			if err2 != nil {
				return 0, err2
			}

			return code, nil
		}
	}

	return 0, errors.New("brightnessctl wrong result")
}
