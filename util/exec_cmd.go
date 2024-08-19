package util

import (
	"os/exec"
	"strings"
)

func ExecCmd(cmd string, arg ...string) (string, error) {
	parts := strings.Fields(cmd)
	parts = append(parts, arg...)
	cmdResult := exec.Command(parts[0], parts[1:]...)

	resultBytes, err := cmdResult.Output()
	if err != nil {
		return "", err
	}

	return string(resultBytes), err
}
