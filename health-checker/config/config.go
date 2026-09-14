package config

import (
	"errors"
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

type Builder struct {
	urls                []string
	healthPath          string
	requestTimeout      time.Duration
	healthCheckInterval time.Duration
	maxConcurrency      int
	errs                []error
}

func NewBuilder() *Builder {
	return &Builder{
		healthPath:          "/health",
		requestTimeout:      5 * time.Second,
		healthCheckInterval: 10 * time.Second,
		maxConcurrency:      10,
	}
}

func (b *Builder) URLs(urls ...string) *Builder {
	b.urls = append(b.urls, urls...)
	return b
}

func (b *Builder) HealthPath(path string) *Builder {
	b.healthPath = path
	return b
}

func (b *Builder) RequestTimeout(d time.Duration) *Builder {
	b.requestTimeout = d
	return b
}

func (b *Builder) HealthCheckInterval(d time.Duration) *Builder {
	b.healthCheckInterval = d
	return b
}

func (b *Builder) MaxConcurrency(n int) *Builder {
	b.maxConcurrency = n
	return b
}

// Build валидирует состояние и возвращает готовый Config.
func (b *Builder) Build() (*Config, error) {
	if len(b.urls) == 0 {
		b.errs = append(b.errs, errors.New("список URL пуст"))
	}

	if b.requestTimeout <= 0 {
		b.errs = append(b.errs, errors.New("requestTimeout должен быть больше 0"))
	}

	if b.healthCheckInterval <= 0 {
		b.errs = append(b.errs, errors.New("healthCheckInterval должен быть больше 0"))
	}

	if b.maxConcurrency <= 0 {
		b.errs = append(b.errs, errors.New("maxConcurrency должен быть больше 0"))

	}
	if len(b.errs) > 0 {
		return nil, errors.Join(b.errs...)
	}

	urlsCopy := make([]string, len(b.urls))
	copy(urlsCopy, b.urls)

	return &Config{
		urls:           urlsCopy,
		healthPath:     b.healthPath,
		requestTimeout: b.requestTimeout,
		maxConcurrency: b.maxConcurrency,
	}, nil
}
