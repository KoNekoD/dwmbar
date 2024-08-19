package keyboard_layout

import (
	"main/util"
	"strings"
)

type Stats struct {
	Lang string
}

func Get() (*Stats, error) {
	output, err := util.ExecCmd("xkblayout-state print %s")
	if err != nil {
		return nil, err
	}

	return &Stats{Lang: strings.ToUpper(output)}, nil
}
