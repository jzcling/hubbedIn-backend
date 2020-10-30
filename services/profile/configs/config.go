package configs

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const configFileName = "env"

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
	v.SetEnvPrefix("api")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	var cfg Config
	if err := v.ReadInConfig(); err != nil {
		return Config{}, errors.Wrap(err, "Failed to read config")
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "Unable to decode into struct")
	}

	return cfg, nil
}
