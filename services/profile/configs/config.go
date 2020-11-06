package configs

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const configFileName = "config"

// Config declares the application configuration variables
type Config struct {
	AppName  string
	Server   serverConfig
	Database dbConfig
}

// dbConfig declares database variables
type dbConfig struct {
	Address    string
	Username   string
	Password   string
	Database   string
	Sslmode    string
	Drivername string
}

// serverConfig declares server variables
type serverConfig struct {
	Address string
	Port    string
}

// LoadConfig load config from file
func LoadConfig() (Config, error) {
	v := viper.New()
	v.SetConfigName(configFileName)
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("../configs")
	v.AddConfigPath("../../configs")

	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	var cfg Config
	if err := v.ReadInConfig(); err != nil {
		return Config{}, errors.Wrap(err, "Failed to read config")
	}

	fmt.Printf("%+v", v)

	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "Unable to decode into struct")
	}

	fmt.Printf("%+v", cfg)

	return cfg, nil
}
