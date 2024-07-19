package config

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

var C *Config

type Config struct {
	DB Database `hcl:"database,block"`

	ListenAddress string `hcl:"listen_address"`
}

func LoadConfig() error {
	var config Config

	if _, err := os.Stat("/etc/incident/config.hcl"); err != nil {
		return fmt.Errorf("config: failed to find config file")
	}
	if err := hclsimple.DecodeFile("/etc/incident/config.hcl", nil, &config); err != nil {
		return fmt.Errorf("config: failed to decode config file: %w", err)
	}

	C = &config

	return nil
}
