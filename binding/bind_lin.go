package binding

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

func BindLin() (string, error) {
	path, errPath := pathFinding()
	if errPath != nil {
		log.Println("Error reading path in bind:", errPath)
		return "", errPath
	}

	cmd := exec.Command("udevadm", "info", "--query=property", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error reading serial number HDD:", err)
		return "", err
	}

	lines := strings.Split(string(output), "\n")
	var serialNumber string
	for _, line := range lines {
		if strings.HasPrefix(line, "ID_SERIAL_SHORT=") {
			serialNumber = strings.TrimPrefix(line, "ID_SERIAL_SHORT=")
			return serialNumber, nil
		}
	}

	return "", errors.New("serial number HDD not found")
}

func pathFinding() (string, error) {
	cmd := exec.Command("df", "/")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error reading root directory: ", err)
		return "", nil
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	if len(lines) > 1 {
		fields := strings.Fields(lines[1])
		if len(fields) >= 6 {
			return fields[0], nil
		}
	}

	return "", errors.New("Path has not been found")
}
