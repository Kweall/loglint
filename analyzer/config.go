package analyzer

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Rules struct {
		Lowercase    bool `yaml:"lowercase"`
		ASCII        bool `yaml:"ascii"`
		SpecialChars bool `yaml:"special_chars"`
		Sensitive    bool `yaml:"sensitive"`
	} `yaml:"rules"`

	SensitiveKeys []string `yaml:"sensitive_keys"`
}

func defaultConfig() *Config {
	return &Config{
		Rules: struct {
			Lowercase    bool `yaml:"lowercase"`
			ASCII        bool `yaml:"ascii"`
			SpecialChars bool `yaml:"special_chars"`
			Sensitive    bool `yaml:"sensitive"`
		}{
			Lowercase:    true,
			ASCII:        true,
			SpecialChars: true,
			Sensitive:    true,
		},
		SensitiveKeys: []string{
			"password",
			"passwd",
			"pwd",
			"token",
			"api_key",
			"apikey",
			"secret",
			"private_key",
			"auth",
		},
	}
}

func loadConfig(path string) *Config {
	cfg := defaultConfig()

	if path == "" {
		return cfg
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg
	}

	if len(cfg.SensitiveKeys) == 0 {
		cfg.SensitiveKeys = defaultConfig().SensitiveKeys
	}

	return cfg
}
