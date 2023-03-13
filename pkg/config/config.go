package config

import (
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TelegramToken string
}

func Init() (*Config, error) {
	var cfg Config
	if err := parseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	os.Setenv("TELEGRAM_BOT_TOKEN", "Your Token")
	if err := viper.BindEnv("telegram_bot_token"); err != nil {
		return err
	}
	cfg.TelegramToken = viper.GetString("telegram_bot_token")

	return nil
}
