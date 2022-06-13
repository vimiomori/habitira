package usr

import (
	"github.com/vimiomori/habitira/cli/config"
)

type Usr interface {
	HasConfig() bool
	Init()
}

type usr struct {
	homeDir string
	osName  string
}

func NewUser(
	homeDir string,
	osName string,
) Usr {
	return &usr{
		homeDir,
		osName,
	}
}

func (u *usr) HasConfig() bool {
	config.Set(u.homeDir)
	err := config.Read()
	if err == nil {
		return true
	} else if config.IsErrNotFound(err) {
		return false
	} else {
		panic("unable to read config file.")
	}
}

// move config creation to config package
// Init asks the user a series of questions
// and creates a config file based on the answers
// Initial setup
func (u *usr) Init() {
	config.CreateFile(u.homeDir)
	// werr := viper.WriteConfig()

	// if werr != nil {
	// 	fmt.Println("failed to write config")
	// 	fmt.Println(werr.Error())
	// }
}
