package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	CONFIG_DIR = "config/habitira"
	CONFIG_NAME = "config"
	CONFIG_TYPE = "yaml"
)

func Set(homeDir string) {
	configDir := fmt.Sprintf("%s/%s", homeDir, CONFIG_DIR)
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

func CreateFile(homeDir string) {
	configPath:= filepath.Join(homeDir, CONFIG_DIR)
	fmt.Printf("Creating config file under %s", configPath)
	err := os.Mkdir(configPath, 0755)
	if err != nil {
		panic(fmt.Errorf("failed to create directory %s: %w", configPath, err))
	}
	filename := filepath.Join(configPath, "config.yaml")
	_, ferr := os.Create(filename)
	if ferr != nil {
		panic(fmt.Sprintf("Failed to create %s", filename))
	}
}
