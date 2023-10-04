package binding

import (
	"errors"
	"os/exec"
	"strings"
	"os"
	"os/user"
)

type BindWin struct{}

func (b *BindWin) ReadDriveNumber() (string, error) {
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

func (b *BindWin) ReadUserName() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func (b *BindWin) ReadMachineName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return hostname, nil
}