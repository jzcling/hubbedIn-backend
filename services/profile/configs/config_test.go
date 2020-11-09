package configs

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	var tests = []struct {
		fileName string
		config   Config
		err      error
	}{
		{"missing", Config{}, errors.New(readErr)},
		{"config_read", Config{}, errors.New(readErr)},
		{
			"config_pass",
			Config{
				AppName: "test",
				Server: ServerConfig{
					Port: "123",
				},
				Database: DbConfig{
					Username: "user",
				},
			},
			nil,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.fileName)
		t.Run(testname, func(t *testing.T) {
			cfg, err := LoadConfig(tt.fileName)
			if cfg != tt.config {
				t.Errorf("got %+v, want %+v", cfg, tt.config)
			}
			if err != nil && !strings.Contains(err.Error(), tt.err.Error()) {
				t.Errorf("got %s, want %s", err, tt.err)
			}
		})
	}
}
