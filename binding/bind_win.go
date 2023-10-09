package binding

import (
	"errors"
	"os/exec"
	"strings"
	"os"
	"os/user"
)

type BindWin struct{
	DriveNumber string `bson:"drivenumber"`
	UserName    string `bson:"username"`
	MachineName string `bson:"machinename"`
}

func (b BindWin) GetDriveNumber() (string, error) {
	cmd := exec.Command("wmic", "diskdrive", "get", "serialnumber")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		serialNumber := strings.TrimSpace(line)
		if serialNumber != "" && serialNumber != "SerialNumber" {
			return strings.TrimRight(serialNumber, "."), nil
		}
	}

	return "", errors.New("serial number HDD not found")
}

func (b BindWin) GetUserName() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	if user.Name != "" {
		return "", errors.New("empty username")
	}
	
	return user.Name, nil
}

func (b BindWin) GetMachineName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	if hostname != "" {
		return "", errors.New("empty hostname")
	}
	return hostname, nil
}