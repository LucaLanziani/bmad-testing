package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg == nil {
		t.Fatal("DefaultConfig returned nil")
	}

	// Default colors depends on TTY, so we just check it's set
	// In tests, it will likely be false since stdout is not a TTY
}

func TestLoadConfigFile_ColorsDisabled(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	configData := map[string]interface{}{
		"output": map[string]interface{}{
			"colors": false,
		},
	}

	data, err := json.Marshal(configData)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load the config
	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if cfg.Output.Colors {
		t.Error("Expected colors to be disabled, but it was enabled")
	}

	if cfg.ColorsEnabled() {
		t.Error("ColorsEnabled() should return false when colors are disabled")
	}
}

func TestLoadConfigFile_ColorsEnabled(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	configData := map[string]interface{}{
		"output": map[string]interface{}{
			"colors": true,
		},
	}

	data, err := json.Marshal(configData)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Load the config
	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if !cfg.Output.Colors {
		t.Error("Expected colors to be enabled, but it was disabled")
	}

	if !cfg.ColorsEnabled() {
		t.Error("ColorsEnabled() should return true when colors are enabled")
	}
}

func TestLoadConfigFile_NonExistent(t *testing.T) {
	// Load config from non-existent file should use defaults
	cfg, err := Load("/nonexistent/config.json")
	if err != nil {
		t.Fatalf("Load should not fail for non-existent file: %v", err)
	}

	if cfg == nil {
		t.Fatal("Config should not be nil")
	}
}

func TestEnvOverride_CONCIGLIA_NO_COLOR(t *testing.T) {
	// Save original env and restore after test
	origValue := os.Getenv("CONCIGLIA_NO_COLOR")
	defer func() {
		if origValue == "" {
			os.Unsetenv("CONCIGLIA_NO_COLOR")
		} else {
			os.Setenv("CONCIGLIA_NO_COLOR", origValue)
		}
	}()

	// Create a config with colors enabled
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	configData := map[string]interface{}{
		"output": map[string]interface{}{
			"colors": true,
		},
	}

	data, err := json.Marshal(configData)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Set CONCIGLIA_NO_COLOR=1
	os.Setenv("CONCIGLIA_NO_COLOR", "1")

	// Load config - should disable colors
	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if cfg.Output.Colors {
		t.Error("CONCIGLIA_NO_COLOR=1 should disable colors")
	}
}

func TestEnvOverride_NO_COLOR(t *testing.T) {
	// Save original env and restore after test
	origValue := os.Getenv("NO_COLOR")
	defer func() {
		if origValue == "" {
			os.Unsetenv("NO_COLOR")
		} else {
			os.Setenv("NO_COLOR", origValue)
		}
	}()

	// Create a config with colors enabled
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	configData := map[string]interface{}{
		"output": map[string]interface{}{
			"colors": true,
		},
	}

	data, err := json.Marshal(configData)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Set NO_COLOR (any value disables colors per standard)
	os.Setenv("NO_COLOR", "1")

	// Load config - should disable colors
	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if cfg.Output.Colors {
		t.Error("NO_COLOR should disable colors")
	}
}

func TestEnvOverride_Priority(t *testing.T) {
	// Save original env and restore after test
	origNO_COLOR := os.Getenv("NO_COLOR")
	origCONCIGLIA := os.Getenv("CONCIGLIA_NO_COLOR")
	defer func() {
		if origNO_COLOR == "" {
			os.Unsetenv("NO_COLOR")
		} else {
			os.Setenv("NO_COLOR", origNO_COLOR)
		}
		if origCONCIGLIA == "" {
			os.Unsetenv("CONCIGLIA_NO_COLOR")
		} else {
			os.Setenv("CONCIGLIA_NO_COLOR", origCONCIGLIA)
		}
	}()

	// Create a config with colors enabled
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	configData := map[string]interface{}{
		"output": map[string]interface{}{
			"colors": true,
		},
	}

	data, err := json.Marshal(configData)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Set both environment variables - NO_COLOR should take priority
	os.Setenv("NO_COLOR", "true")
	os.Setenv("CONCIGLIA_NO_COLOR", "0") // This should be ignored because NO_COLOR is set

	// Load config - should disable colors due to NO_COLOR
	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if cfg.Output.Colors {
		t.Error("NO_COLOR should take priority and disable colors")
	}
}

func TestEnvOverride_CONCIGLIA_NO_COLOR_NotOne(t *testing.T) {
	// Save original env and restore after test
	origValue := os.Getenv("CONCIGLIA_NO_COLOR")
	defer func() {
		if origValue == "" {
			os.Unsetenv("CONCIGLIA_NO_COLOR")
		} else {
			os.Setenv("CONCIGLIA_NO_COLOR", origValue)
		}
	}()

	// Create a config with colors enabled
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.json")

	configData := map[string]interface{}{
		"output": map[string]interface{}{
			"colors": true,
		},
	}

	data, err := json.Marshal(configData)
	if err != nil {
		t.Fatalf("Failed to marshal config: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		t.Fatalf("Failed to write config file: %v", err)
	}

	// Set CONCIGLIA_NO_COLOR to something other than "1"
	os.Setenv("CONCIGLIA_NO_COLOR", "0")

	// Load config - should NOT disable colors because value is not "1"
	cfg, err := Load(configPath)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if !cfg.Output.Colors {
		t.Error("CONCIGLIA_NO_COLOR=0 should not disable colors (only CONCIGLIA_NO_COLOR=1)")
	}
}
