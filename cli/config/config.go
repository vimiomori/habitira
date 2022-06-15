package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/vimiomori/habitira/cli/env"
)

const (
	CONFIG_DIR  = ".config/habitira"
	CONFIG_FILE = "config.yaml"
	CONFIG_NAME = "config"
	CONFIG_TYPE = "yaml"
)

type Config interface {
	Setup()
	Read() error
	IsErrNotFound(err error) bool
	CreateFile()
	GetHomeDir() string
	GetUserName() string
}

type config struct {
	homeDir        string
	userName     string
}

func NewConfig() Config {
	return &config{
		env.GetOSHomeDir(),
		env.GetOSUserName(),
	}
}

func (c *config) Setup() {
	configDir := fmt.Sprintf("%s/%s", c.homeDir, CONFIG_DIR)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
}

func (c *config) Read() error {
	return viper.ReadInConfig()
}

func (c *config) IsErrNotFound(err error) bool {
	_, ok := err.(viper.ConfigFileNotFoundError)
	return ok
}

func (c *config) CreateFile() {
	configPath := filepath.Join(c.homeDir, CONFIG_DIR)
	fmt.Printf("Creating config file under %s", configPath)
	if err := os.MkdirAll(configPath, 0755); err != nil {
		panic(fmt.Errorf("failed to create directory %s: %w", configPath, err))
	}
	filename := filepath.Join(configPath, CONFIG_FILE)
	_, ferr := os.Create(filename)
	if ferr != nil {
		panic(fmt.Sprintf("Failed to create %s", filename))
	}
}

func (c *config) GetHomeDir() string {
	return c.homeDir
}

func (c *config) GetUserName() string {
	name := viper.GetString("username")
	if name != "" {
		return name
	}
	return c.userName
}
