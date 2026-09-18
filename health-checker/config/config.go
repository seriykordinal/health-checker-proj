package config

import (
	"time"
)

type Config struct {
	urls                []string
	healthPath          string
	requestTimeout      time.Duration
	healthCheckInterval time.Duration
	maxConcurrency      int
}

func (c *Config) URLs() []string                     { return c.urls }
func (c *Config) HealthPath() string                 { return c.healthPath }
func (c *Config) RequestTimeout() time.Duration      { return c.requestTimeout }
func (c *Config) MaxConcurrency() int                { return c.maxConcurrency }
func (c *Config) HealthCheckInterval() time.Duration { return c.healthCheckInterval }
