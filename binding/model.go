package binding

type BindHW struct {
	DriveNumber string `bson:"drivenumber"`
	UserName    string `bson:"username"`
	MachineName string `bson:"machinename"`
}

type Binder interface {
	ReadDriveNumber() (string, error)
	ReadUserName() (string, error)
	ReadMachineName() (string, error)
}

func (b *BindHW) SetValues(binder Binder) error {
	number, err := binder.ReadDriveNumber()
	if err != nil {
		return err
	}

	b.DriveNumber = number

	usernamePC, err := binder.ReadUserName()
	if err != nil {
		return err
	}

	b.UserName = usernamePC

	namePC, err := binder.ReadMachineName()
	if err != nil {
		return err
	}

	b.MachineName = namePC

	return nil
}
