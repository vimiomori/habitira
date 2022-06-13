package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	CONFIG_DIR = "./config/habitira"
	CONFIG_NAME = "config"
	CONFIG_TYPE = "yaml"
)

func Set() {
	homeDir, _ := os.UserHomeDir()
	configDir := fmt.Sprintf("%s/.config/habitira", homeDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
}

func Read() error {
	return viper.ReadInConfig()
}

func IsErrNotFound(err error) bool {
	_, ok := err.(viper.ConfigFileNotFoundError)
	return ok
}
