package usr

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	CONFIG_DIR = "./config/habitira"
	CONFIG_NAME = "config"
	CONFIG_TYPE = "yaml"
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
	usr, _ := user.Current()
	homeDir, _ := os.UserHomeDir()
	configDir := fmt.Sprintf("%s/.config/habitira", homeDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	err := viper.ReadInConfig()
	if err == nil {
		return true
	}

	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		fmt.Printf("welcome %s", usr.Name)
		viper.Set("name", "meeeooo")
		path := filepath.Join(configDir)
		err := os.Mkdir(path, 0755)
		if err != nil {
			panic(fmt.Errorf("failed to create directory %s: %w", path, err))
		}
		filename := filepath.Join(configDir, "config.yaml")
		_, ferr := os.Create(filename)
		if ferr != nil {
				panic(fmt.Sprintf("Failed to create %s", filename))
		}

		werr := viper.WriteConfig() 

		if werr != nil {
			fmt.Println("failed to write config")
			fmt.Println(werr.Error())
		}
	} else {
		fmt.Println(err.Error())
		panic(fmt.Sprintf("Failed to read config from %s", configDir))
	}

	return false
}

func (u *usr) CreateConfig() {}

