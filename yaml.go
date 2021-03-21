package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func LoadFile(path string) (*Configuration, error) {
	var (
		raw []byte
		err error
	)
	if raw, err = ioutil.ReadFile(path); err != nil {
		return nil, fmt.Errorf("configuration read error: %w", err)
	}
	var cfg Configuration
	if err = yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("configuration parsing error: %w", err)
	}
	return &cfg, nil
}
