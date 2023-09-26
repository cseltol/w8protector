package macExtract

import (
	"errors"
	"log"
	"net"
)

func ExtractMac() (string, error) {
	var activeMac string

	netInterfaces, err := net.Interfaces()
	if err != nil {
		log.Printf("failed to get interfaces: %v", err)
		return "", err
	}

	for _, activeInterfase := range netInterfaces {
		if activeInterfase.Flags&net.FlagUp != 0 && activeInterfase.Flags&net.FlagLoopback == 0 {
			activeMac = activeInterfase.HardwareAddr.String()

			log.Printf("interface name: %s", activeInterfase.Name)
			log.Printf("MAC address %s", activeMac)
		}
	}

	if activeMac == "" {
		return "", errors.New("no active interface found")
	}

	return activeMac, nil
}
