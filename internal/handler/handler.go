package handler

import (
	"github.com/edwin-Marrima/Pod-net-route-guard/internal/schema"
	"gopkg.in/yaml.v3"
	"os"
)

func readConfiguration(configFilePath string) (*schema.Config, error) {
	cfg, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	var config schema.Config
	err = yaml.Unmarshal(cfg, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func Apply() error {
	// read configuration
	return nil
}
