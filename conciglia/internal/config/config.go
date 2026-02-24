package config

import (
	"encoding/json"
	"golang.org/x/term"
	"os"
)

// OutputConfig holds output-related configuration
type OutputConfig struct {
	Colors bool `json:"colors"`
}

// Config represents the complete Conciglia configuration
type Config struct {
	Output OutputConfig `json:"output"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Output: OutputConfig{
			Colors: isTerminal(),
		},
	}
}

// Load loads configuration from a file and applies environment variable overrides
func Load(configPath string) (*Config, error) {
	cfg := DefaultConfig()

	// Load from file if it exists
	if configPath != "" {
		data, err := os.ReadFile(configPath)
		if err != nil {
			if !os.IsNotExist(err) {
				return nil, err
			}
			// File doesn't exist, use defaults
		} else {
			if err := json.Unmarshal(data, cfg); err != nil {
				return nil, err
			}
		}
	}

	// Apply environment variable overrides
	cfg.applyEnvOverrides()

	return cfg, nil
}

// applyEnvOverrides applies environment variable overrides to the config
func (c *Config) applyEnvOverrides() {
	// Check NO_COLOR environment variable (de facto standard)
	if os.Getenv("NO_COLOR") != "" {
		c.Output.Colors = false
		return
	}

	// Check CONCIGLIA_NO_COLOR environment variable (app-specific override)
	if os.Getenv("CONCIGLIA_NO_COLOR") == "1" {
		c.Output.Colors = false
		return
	}
}

// isTerminal checks if stdout is connected to a terminal
func isTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// ColorsEnabled returns whether colors should be used in output
func (c *Config) ColorsEnabled() bool {
	return c.Output.Colors
}
