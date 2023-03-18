package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	TelegramBotToken string
	DatabaseURL      string
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		viper.AutomaticEnv()
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := fromEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	return nil
}

func fromEnv(cfg *Config) error {
	if err := viper.BindEnv("TELEGRAM_BOT_TOKEN"); err != nil {
		return err
	}
	cfg.TelegramBotToken = viper.GetString("TELEGRAM_BOT_TOKEN")

	if err := viper.BindEnv("DB_USER"); err != nil {
		return err
	}
	if err := viper.BindEnv("DB_PASSWORD"); err != nil {
		return err
	}
	if err := viper.BindEnv("DB_HOST"); err != nil {
		return err
	}
	if err := viper.BindEnv("DB_NAME"); err != nil {
		return err
	}

	cfg.DatabaseURL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		viper.GetString("DB_USER"), viper.GetString("DB_PASSWORD"), viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"))
	return nil
}

func setUpViper() error {
	viper.SetConfigFile(".env")

	return viper.ReadInConfig()
}
