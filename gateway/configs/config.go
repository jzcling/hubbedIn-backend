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

	readErr         string = "Failed to read config"
	unmarshalErr    string = "Unable to decode into struct"
	mapstructureErr string = "Error decoding mapstructure"
)

// Config declares the application configuration variables
type Config struct {
	AppName string       `mapstructure:"appname"`
	Server  ServerConfig `mapstructure:",squash"`
}

// ServerConfig declares server variables
type ServerConfig struct {
	Address string `mapstructure:"server_address"`
	Port    string `mapstructure:"server_port"`
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
