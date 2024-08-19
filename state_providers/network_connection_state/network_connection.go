package network_connection_state

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"main/util"
	"os"
	"strings"
)

// Get network_connection statistics
func Get() (*Stats, error) {
	path := "/sys/class/net/"
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	wiredConnected, wirelessConnected := false, false
	wirelessConnectionName := ""

	wiredInterfaceName, wirelessInterfaceName := "", ""
	for i := range dirEntries {
		name := dirEntries[i].Name()

		hasPrefix, isUp, err2 := checkOperstateByPrefix(
			"e",
			name,
			path,
		)

		// Force stop
		if err2 != nil {
			return nil, err2
		}

		if hasPrefix {
			if isUp {
				wiredConnected = true
			}
			wiredInterfaceName = name
			continue
		}

		hasPrefix, isUp, err = checkOperstateByPrefix(
			"wl",
			name,
			path,
		)
		if err != nil {
			return nil, err
		}
		if hasPrefix {
			if isUp {
				wirelessConnected = true
				wirelessInterfaceName = name
				wirelessConnectionNamee, err44 := getWirelessConnectionName(
					wirelessInterfaceName,
				)
				if err44 == nil {
					wirelessConnectionName = wirelessConnectionNamee
				}
			}
			continue
		}
	}

	return &Stats{
			WiredConnected:         wiredConnected,
			WirelessConnected:      wirelessConnected,
			WirelessConnectionName: wirelessConnectionName,
			WiredInterfaceName:     wiredInterfaceName,
			WirelessInterfaceName:  wirelessInterfaceName,
		},
		nil
}

func getWirelessConnectionName(iface string) (string, error) {
	output, err := util.ExecCmd("iwctl station", iface, "show")
	if err != nil {
		return "", err
	}

	lines := strings.Split(output, "\n")

	for _, line := range lines {
		connectionString, found := strings.CutPrefix(
			strings.TrimSpace(line),
			"Connected network",
		)
		if found {
			connectionString = strings.TrimSpace(connectionString)
			return connectionString, nil
		}
	}

	return "", errors.New("iwctl wrong result")
}

// Stats represents network_connection info
type Stats struct {
	WiredConnected         bool
	WirelessConnected      bool
	WirelessConnectionName string

	WiredInterfaceName    string
	WirelessInterfaceName string
}

func (s *Stats) IsConnected() bool {
	return s.WiredConnected || s.WirelessConnected
}

func (s *Stats) IsNotConnected() bool {
	return !s.IsConnected()
}

func (s *Stats) GetActiveInterfaceName() (string, error) {
	if s.WiredConnected {
		return s.WiredInterfaceName, nil
	}

	if s.WirelessConnected {
		return s.WirelessInterfaceName, nil
	}

	return "", errors.New("all network is off")
}

type OperstateError struct {
	message string
}

func (e *OperstateError) Error() string {
	return e.message
}

func checkOperstateByPrefix(prefix, dirname, path string) (bool, bool, error) {
	hasPrefix := strings.HasPrefix(dirname, prefix)
	if hasPrefix {
		file, err := os.Open(path + "/" + dirname + "/operstate")
		if err != nil {
			return true, false, err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				_, err2 := fmt.Fprintf(os.Stderr, "%s\n", err)
				if err2 != nil {
					return
				}
				return
			}
		}(file)
		isUp, _ := getOperstateStatus(file)
		if isUp {
			return true, true, nil
		}
	}
	return false, false, nil
}

func getOperstateStatus(out io.Reader) (bool, error) {
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		switch scanner.Text() {
		case "up":
			return true, nil
		case "down":
			return false, nil
		default:
			err := &OperstateError{
				message: "Wrong Operstate file text",
			}

			return false, err
		}
	}
	return false, &OperstateError{message: "Scanner problem"}
}
