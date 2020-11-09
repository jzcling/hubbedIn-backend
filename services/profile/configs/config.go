package configs

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Config declares the application configuration variables
type Config struct {
	AppName  string       `mapstructure:"appname"`
	Server   serverConfig `mapstructure:",squash"`
	Database dbConfig     `mapstructure:",squash"`
}

// dbConfig declares database variables
type dbConfig struct {
	Address    string `mapstructure:"database_address"`
	Username   string `mapstructure:"database_username"`
	Password   string `mapstructure:"database_password"`
	Database   string `mapstructure:"database_database"`
	Sslmode    string `mapstructure:"database_sslmode"`
	Drivername string `mapstructure:"database_drivername"`
}

// serverConfig declares server variables
type serverConfig struct {
	Address string `mapstructure:"server_address"`
	Port    string `mapstructure:"server_port"`
}

// LoadConfig load config from file
func LoadConfig() (Config, error) {
	var result map[string]interface{}
	var cfg Config

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("../configs")
	v.AddConfigPath("../../configs")

	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return Config{}, errors.Wrap(err, "Failed to read config")
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, "Unable to decode into struct")
	}

	err = mapstructure.Decode(result, &cfg)
	if err != nil {
		fmt.Println("error decoding")
	}

	return cfg, nil
}
