package config

import "testing"

func TestBasicExample_Load(t *testing.T) {
	_, err := LoadFile("testdata/basic-example.yaml")

	if err != nil {
		t.Fatalf("success parsing expected but got error: %s", err)
	}
	//	t.Logf("%+v", cfg)
}

func TestBasicExample_ServicesNames(t *testing.T) {
	cfg, _ := LoadFile("testdata/basic-example.yaml")

	if len(cfg.Services) != 2 {
		t.Fatalf("2 services expected but got %d", len(cfg.Services))
	}
}
