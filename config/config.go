package config

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

// Config holds the global configuration for the library.
type Config struct {
	mu             sync.RWMutex
	GlobalMetadata map[string]interface{}
	LogLevel       string
	Environment    string
	ServiceName    string
	Region         string
}

// globalConfig is a singleton instance of Config.
var globalConfig *Config
var once sync.Once

// GetConfig initializes or returns the singleton Config instance.
func GetConfig() *Config {
	once.Do(func() {
		globalConfig = &Config{
			GlobalMetadata: map[string]interface{}{
				"version":     "1.0.0",
				"serviceName": "FlexiResponseGo",
				"region":      "default-region",
			},
			LogLevel:    "info",
			Environment: "production",
			ServiceName: "FlexiResponseGo",
			Region:      "default-region",
		}
	})
	return globalConfig
}

// UpdateMetadata updates global metadata key-value pairs dynamically.
func (c *Config) UpdateMetadata(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.GlobalMetadata[key] = value
}

// GetMetadata retrieves the value for a metadata key.
func (c *Config) GetMetadata(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.GlobalMetadata[key]
	return value, exists
}

// UpdateLogLevel dynamically updates the logging level.
func (c *Config) UpdateLogLevel(level string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.LogLevel = level
}

// UpdateEnvironment dynamically updates the environment.
func (c *Config) UpdateEnvironment(env string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Environment = env
}

// LoadFromFile loads configuration from a JSON file.
func (c *Config) LoadFromFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileConfig Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&fileConfig); err != nil {
		return err
	}

	// Lock and update the current configuration
	c.mu.Lock()
	defer c.mu.Unlock()

	c.GlobalMetadata = fileConfig.GlobalMetadata
	c.LogLevel = fileConfig.LogLevel
	c.Environment = fileConfig.Environment
	c.ServiceName = fileConfig.ServiceName
	c.Region = fileConfig.Region

	return nil
}

// SaveToFile saves the current configuration to a JSON file.
func (c *Config) SaveToFile(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	c.mu.RLock()
	defer c.mu.RUnlock()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(c); err != nil {
		return err
	}

	return nil
}

// ReloadFromFile reloads configuration from the given file dynamically.
func (c *Config) ReloadFromFile(filepath string) error {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return errors.New("configuration file not found")
	}
	return c.LoadFromFile(filepath)
}
