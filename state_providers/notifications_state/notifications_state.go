package notifications_state

import (
	"main/util"
	"strings"
)

type Stats struct {
	IsDisabled bool
}

func Get() (*Stats, error) {
	output, err := util.ExecCmd("dunstctl is-paused")
	if err != nil {
		return nil, err
	}

	output = strings.TrimSpace(output)

	isDisabled := strings.ToUpper(output) == "TRUE"

	return &Stats{IsDisabled: isDisabled}, nil
}
