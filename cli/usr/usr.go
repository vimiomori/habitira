package usr

import (
	"fmt"

	"github.com/vimiomori/habitira/cli/config"
)

type Usr interface {
	SetupConfig()
	HasConfig() bool
	Init()
	WelcomeBack()
}

type usr struct {
	config config.Config
}

func NewUser(
) Usr {
	return &usr{
		config.NewConfig(),
	}
}

func (u *usr) SetupConfig() {
	u.config.Setup()
}

func (u *usr) HasConfig() bool {
	err := u.config.Read()
	if err == nil {
		return true
	} else if u.config.IsErrNotFound(err) {
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
	fmt.Printf("Welcome %s!!\n", u.config.GetUserName())
	// sleep
	fmt.Println("Let's begin your initial setup.")
	u.config.CreateFile()
	// werr := viper.WriteConfig()

	// if werr != nil {
	// 	fmt.Println("failed to write config")
	// 	fmt.Println(werr.Error())
	// }
}

// TODO: rename
// display options for view/edit
func (u *usr) WelcomeBack() {
	fmt.Printf("Welcome back, %s", u.config.GetUserName())
}
