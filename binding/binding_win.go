package binding

import (
	"errors"
	"os/exec"
	"strings"
)

func BindWin() (string, error) {
	cmd := exec.Command("wmic", "diskdrive", "get", "serialnumber")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		serialNumber := strings.TrimSpace(line)
		if serialNumber != "" && serialNumber != "SerialNumber" {
			serialNumber = strings.TrimRight(serialNumber, ".")
			return serialNumber, nil
		}
	}

	return "", errors.New("Serial number HDD not found")
}
