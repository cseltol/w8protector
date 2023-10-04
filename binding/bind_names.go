package binding

import (
	"os"
	"os/user"
)

func ReadUsername() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return user.Name, nil
}

func BindMachineName() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return hostname, nil
}
