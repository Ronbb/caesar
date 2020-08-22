package config

import (
	"fmt"

	"go.uber.org/config"
)

var globalConfig Config

// Global Get global config.
func Global() *Config {
  return &globalConfig
}

// HTTPServerConfig HTTP server Config.
type HTTPServerConfig struct {
  baseServerConfig
  Host string
  Port int
}

// HTTPSServerConfig HTTPS server Config.
type HTTPSServerConfig struct {
  baseServerConfig
  Host string
  Port int
  Fallback *HTTPSServerFallback
  Domain string
  Email string
  Enabled bool
  AutoCert bool
}

// HTTPSServerFallback HTTPS server fallback.
type HTTPSServerFallback struct {
  baseServerConfig
  Host string
  Port int
  Domain string
}


type baseServerConfig interface {
  Address() string
}

// Config Global config.
type Config struct {
  Name string
  Debug bool
  HTTP *HTTPServerConfig
  HTTPS *HTTPSServerConfig
}

func init() {
  defer validate()

	file := config.File(*ConfigFile)
	yaml, err := config.NewYAML(file)
	if err != nil {
		panic(err)
	}

	root := yaml.Get(config.Root)

  root.Populate(&globalConfig)
}

func validate()  {
  // TODO
}

// Address Server address.
func (c *HTTPServerConfig) Address() string {
  return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Address Server address.
func (c *HTTPSServerConfig) Address() string {
  return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Address Fallback address.
func (c *HTTPSServerFallback) Address() string {
  return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
