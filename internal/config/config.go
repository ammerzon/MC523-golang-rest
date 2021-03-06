package config

import (
  "os"
  "sync"
)

var config *Config
var once sync.Once

// A Config represents configurable parameters of the program.
type Config struct {
  DatabaseHost     string
  DatabaseName     string
  DatabasePassword string
  DatabaseUsername string
}

// GetConfig loads the configurable parameters through environment variables.
// The values will only be loaded once and therefore you need to restart the program if you changed your environment.
func GetConfig() *Config {
  once.Do(func() {
    config = &Config{
      DatabaseHost:     getString("APP_DB_HOST", "localhost"),
      DatabaseName:     getString("APP_DB_NAME", "productdb"),
      DatabasePassword: getString("APP_DB_PASSWORD", "postgres"),
      DatabaseUsername: getString("APP_DB_USERNAME", "postgres"),
    }
  })
  return config
}

// getString returns the value of the environment variable.
// If it does not exist the default value will be returned.
func getString(key string, defaultVal string) string {
  if value, exists := os.LookupEnv(key); exists && value != "" {
    return value
  }

  return defaultVal
}
