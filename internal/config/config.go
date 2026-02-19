package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Verbose    bool   `mapstructure:"verbose"`
	NoColor    bool   `mapstructure:"no_color"`
	AI         bool   `mapstructure:"ai"`
	AIProvider string `mapstructure:"ai_provider"`
	Language   string `mapstructure:"language"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/runex")
	viper.AddConfigPath(".")

	viper.SetDefault("verbose", false)
	viper.SetDefault("no_color", false)
	viper.SetDefault("ai", false)
	viper.SetDefault("ai_provider", "openai")
	viper.SetDefault("language", "")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if envVerbose := os.Getenv("RUNEX_VERBOSE"); envVerbose != "" {
		cfg.Verbose = envVerbose == "true" || envVerbose == "1"
	}
	if envNoColor := os.Getenv("RUNEX_NO_COLOR"); envNoColor != "" {
		cfg.NoColor = envNoColor == "true" || envNoColor == "1"
	}
	if envAI := os.Getenv("RUNEX_AI"); envAI != "" {
		cfg.AI = envAI == "true" || envAI == "1"
	}

	return &cfg, nil
}

func (c *Config) SetVerbose(v bool)    { c.Verbose = v }
func (c *Config) SetNoColor(v bool)    { c.NoColor = v }
func (c *Config) SetAI(v bool)         { c.AI = v }
func (c *Config) SetLanguage(l string) { c.Language = l }
