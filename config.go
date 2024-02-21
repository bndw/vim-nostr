package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Relays     map[string]relayConfig `json:"relays"`
	PrivateKey string                 `json:"privatekey"`
}

func (c Config) WriteRelays() []string {
	var urls []string

	for url, config := range c.Relays {
		if !config.Write {
			continue
		}
		urls = append(urls, url)
	}

	return urls
}

type relayConfig struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

func loadConfig(config *Config) error {
	var configDir string
	switch runtime.GOOS {
	case "darwin":
		configDir, _ = os.UserHomeDir()
		configDir = filepath.Join(configDir, ".config")
	default:
		configDir, _ = os.UserConfigDir()
	}

	configPath := filepath.Join(configDir, "nostr", "config.json")
	b, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, config)
}
