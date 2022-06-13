package usr

import (
	"fmt"

	"github.com/vimiomori/habitira/cli/config"
)

type Usr interface {
	HasConfig() bool
	CreateConfig()
}

type usr struct {
}

func NewUser() Usr {
	return &usr{}
}

func (u *usr) HasConfig() bool {
	config.Set()
	err := config.Read()
	if err == nil {
		return true
	} else if config.IsErrNotFound(err) {
		return false
	} else {
		panic(fmt.Sprintf("unable to read config file."))
	}
}

// move config creation to config package
// CreateConfig asks the user a series of questions
// and creates a config file based on the answers
func (u *usr) CreateConfig() {
	// usr, _ := user.Current()
	// fmt.Printf("welcome %s", usr.Name)
	// viper.Set("name", "meeeooo")
	// path := filepath.Join(configDir)
	// err := os.Mkdir(path, 0755)
	// if err != nil {
	// 	panic(fmt.Errorf("failed to create directory %s: %w", path, err))
	// }
	// filename := filepath.Join(configDir, "config.yaml")
	// _, ferr := os.Create(filename)
	// if ferr != nil {
	// 		panic(fmt.Sprintf("Failed to create %s", filename))
	// }

	// werr := viper.WriteConfig() 

	// if werr != nil {
	// 	fmt.Println("failed to write config")
	// 	fmt.Println(werr.Error())
	// }
}

