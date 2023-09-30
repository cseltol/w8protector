package binding

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

func BindMac() (string, error) {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error when executing the command: ", err)
		return "", nil
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Serial Number (system)") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				serialNumber := strings.TrimSpace(parts[1])

				return serialNumber, nil
			}
		}
	}

	return "", errors.New("Serial number HDD not found")
}
