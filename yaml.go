package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

func LoadComposeFile(path string) (*Compose, error) {
	var (
		raw []byte
		err error
	)
	if raw, err = ioutil.ReadFile(path); err != nil {
		return nil, fmt.Errorf("compose file read error: %w", err)
	}
	var cfg Compose
	if err = yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("compose file parsing error: %w", err)
	}
	return &cfg, nil
}

func LoadFile(path string) (*Extended, error) {
	var (
		raw []byte
		err error
	)
	if raw, err = ioutil.ReadFile(path); err != nil {
		return nil, fmt.Errorf("configuration read error: %w", err)
	}
	var cfg Extended
	if err = yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("configuration parsing error: %w", err)
	}
	return &cfg, nil
}
