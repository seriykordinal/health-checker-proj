package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type yamlConfig struct {
	URLs           []string `yaml:"urls"`
	HealthPath     string   `yaml:"health_path"`
	RequestTimeout string   `yaml:"request_timeout"` // "5s", "1m" и т.п.
	MaxConcurrency int      `yaml:"max_concurrency"`
}

func LoadFromYAML(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать %q: %w", path, err)
	}

	var raw yamlConfig
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("не удалось распарсить YAML: %w", err)
	}

	b := NewBuilder().URLs(raw.URLs...)

	if raw.HealthPath != "" {
		b.HealthPath(raw.HealthPath)
	}

	if raw.RequestTimeout != "" {
		d, err := time.ParseDuration(raw.RequestTimeout)
		if err != nil {
			return nil, fmt.Errorf("некорректный request_timeout: %w", err)
		}
		b.RequestTimeout(d)
	}

	if raw.MaxConcurrency > 0 {
		b.MaxConcurrency(raw.MaxConcurrency)
	}

	return b.Build()
}
