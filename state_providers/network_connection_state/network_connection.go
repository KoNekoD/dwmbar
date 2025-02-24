package network_connection_state

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"main/util"
	"os"
	"os/exec"
	"strings"
)

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
				nameStr, errConn := getWirelessConnectionName(name)
				if errConn == nil {
					wirelessConnectionName = nameStr
				} else {
					wirelessInterfaceName = name
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
	if commandExists("iwctl") {
		return getWirelessConnectionNameFromIwctl(iface)
	} else if commandExists("wpa_cli") {
		return getWirelessConnectionNameFromWPACli(iface)
	} else if commandExists("nmcli") {
		return getWirelessConnectionNameFromNmcli()
	}
	return "", errors.New("no supported software found to retrieve wireless connection information")
}

func commandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func getWirelessConnectionNameFromIwctl(iface string) (string, error) {
	output, err := util.ExecCmd("iwctl", "station", iface, "show")
	if err != nil {
		return "", err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "Connected network") {
			connectionString := strings.TrimSpace(strings.TrimPrefix(trimmed, "Connected network"))
			return connectionString, nil
		}
	}

	return "", errors.New("iwctl: failed to determine network name")
}

func getWirelessConnectionNameFromWPACli(iface string) (string, error) {
	output, err := util.ExecCmd("wpa_cli", "-i", iface, "status")
	if err != nil {
		return "", err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ssid=") {
			return strings.TrimSpace(strings.TrimPrefix(line, "ssid=")), nil
		}
	}

	return "", errors.New("wpa_cli: failed to determine network name")
}

func getWirelessConnectionNameFromNmcli() (string, error) {
	output, err := util.ExecCmd("nmcli", "-t", "-f", "active,ssid", "dev", "wifi")
	if err != nil {
		return "", err
	}

	lines := strings.Split(output, "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 && parts[0] == "yes" {
			return parts[1], nil
		}
	}

	return "", errors.New("nmcli: failed to determine active connection")
}

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

	return "", errors.New("all network interfaces are disconnected")
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
			return false, &OperstateError{
				message: "Invalid operstate file content",
			}
		}
	}
	return false, &OperstateError{message: "Scanner issue"}
}
