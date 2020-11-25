package configs

import (
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	// FileName declares file name for config file
	FileName string = "config"
	// TestFileName declares file name for test config file
	TestFileName string = "config_test"

	readErr         string = "Failed to read config"
	unmarshalErr    string = "Unable to decode into struct"
	mapstructureErr string = "Error decoding mapstructure"
)

// Config declares the application configuration variables
type Config struct {
	AppName  string       `mapstructure:"appname"`
	Server   ServerConfig `mapstructure:",squash"`
	Database DbConfig     `mapstructure:",squash"`
}

// ServerConfig declares server variables
type ServerConfig struct {
	Address string `mapstructure:"server_address"`
	Port    string `mapstructure:"server_port"`
}

// DbConfig declares database variables
type DbConfig struct {
	Address    string `mapstructure:"database_address"`
	Port       string `mapstructure:"database_port"`
	Username   string `mapstructure:"database_username"`
	Password   string `mapstructure:"database_password"`
	Database   string `mapstructure:"database_database"`
	Sslmode    string `mapstructure:"database_sslmode"`
	Drivername string `mapstructure:"database_drivername"`
}

// LoadConfig load config from file
func LoadConfig(fileName string) (Config, error) {
	var result map[string]interface{}
	var cfg Config

	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType("env")
	v.AddConfigPath(".")
	v.AddConfigPath("../configs")
	v.AddConfigPath("../../configs")
	v.AddConfigPath("../tests")

	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return Config{}, errors.Wrap(err, readErr)
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, unmarshalErr)
	}

	err = mapstructure.Decode(result, &cfg)
	if err != nil {
		return Config{}, errors.Wrap(err, mapstructureErr)
	}

	return cfg, nil
}
