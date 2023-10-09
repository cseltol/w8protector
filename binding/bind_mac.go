package binding

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type BindMac struct{
	DriveNumber string `bson:"drivenumber"`
	UserName    string `bson:"username"`
	MachineName string `bson:"machinename"`
}

func (b BindMac) GetDriveNumber() (string, error) {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		return "", errors.New("Serial number HDD not found" + err.Error())
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Serial Number (system)") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1]), nil
			}
		}
	}

	return "", errors.New("serial number HDD not found")
}

func (b BindMac) GetUserName() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	if user.Name != "" {
		return "", errors.New("empty username")
	}
	
	return user.Name, nil
}

func (b BindMac) GetMachineName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	if hostname != "" {
		return "", errors.New("empty hostname")
	}
	return hostname, nil
}
