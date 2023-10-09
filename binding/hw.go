package binding

import (
	"errors"
	"runtime"
)

type BindHW struct {
	DriveNumber string `bson:"drivenumber"`
	UserName    string `bson:"username"`
	MachineName string `bson:"machinename"`
}

type Binder interface {
	GetDriveNumber() (string, error) 
	GetMachineName() (string, error) 
	GetUserName() (string, error) 
}

const (
	WIN_OS = "windows"
	MAC_OS = "darwin"
	UNIX_OS = "unix"
)

func getHWBind() BindHW {
	var empty BindHW
	drive, machine, name := getBinded()
	if drive == "" {
		return empty
	}
	if machine == "" {
		return empty
	}
	if name == "" {
		return empty
	}

	return BindHW{
		DriveNumber: drive,
		MachineName: machine,
		UserName: name,
	}
}

func getBinded() (string, string, string) {
	var err error
	var drive, machine, name string

	binder := getBinder()

	if drive, err = binder.GetDriveNumber(); err != nil {
		return "", "", ""
	}

	if machine, err = binder.GetMachineName(); err != nil {
		return "", "", ""
	}

	if name, err = binder.GetUserName(); err != nil {
		return "", "", ""
	} 
	return drive, machine, name
}

func getBinder() Binder {
	var binder Binder
	switch runtime.GOOS {
	case WIN_OS:
		binder = BindWin{}
	case MAC_OS:
		binder = BindMac{}
	case UNIX_OS:
		binder = BindLin{}
	}
	return binder
}

func BindMachine() error {
	hw := getHWBind()
	if hw == (BindHW{}) {
		return errors.New("empty hardware to bind")
	}
	/*  !TODO:
		Some binding logic here...
	*/

	return nil
}