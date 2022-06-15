package env

import (
	"os"
	u "os/user"
)

func GetOSHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("failed to retrieve user home directory")
	}

	return homeDir
}

func GetOSUserName() string {
	osUser, err := u.Current()
	if err != nil {
		panic("failed to retrieve the current user of the OS")
	}

	return osUser.Name
}
