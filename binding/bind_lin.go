package binding

import (
	"errors"
	"os/exec"
	"strings"
	"os"
	"os/user"
)

type BindLin struct {
	DriveNumber string `bson:"drivenumber"`
	UserName    string `bson:"username"`
	MachineName string `bson:"machinename"`
}

func (b BindLin) GetDriveNumber() (string, error)  {
	path, errPath := pathFinding()
	if errPath != nil {
		return "", errors.New("Error reading path in bind:" + errPath.Error())
	}

	cmd := exec.Command("udevadm", "info", "--query=property", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New("Error reading serial number HDD:" + err.Error())
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "ID_SERIAL_SHORT=") {
			return strings.TrimPrefix(line, "ID_SERIAL_SHORT="), nil
		}
	}
	
	return "", errors.New("serial number HDD not found")
}

func pathFinding() (string, error) {
	cmd := exec.Command("df", "/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New("error reading root directory: " + err.Error())
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	if len(lines) < 1 {
		return "", errors.New("failed to get path")
	}

	fields := strings.Fields(lines[1])
	if len(fields) < 6 {
		return "", errors.New("path has not been found")
	}

	return fields[0], nil
}


func (b BindLin) GetUserName() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	if user.Name != "" {
		return "", errors.New("empty username")
	}
	
	return user.Name, nil
}

func (b BindLin) GetMachineName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	if hostname != "" {
		return "", errors.New("empty hostname")
	}
	return hostname, nil
}