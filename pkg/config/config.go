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
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := loadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadEnv(cfg *Config) error {
	if err := viper.BindEnv("TELEGRAM_BOT_TOKEN"); err != nil {
		return err
	}
	cfg.TelegramBotToken = viper.GetString("TELEGRAM_BOT_TOKEN")

	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbName := viper.GetString("DB_NAME")
	cfg.DatabaseURL = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbName)

	return nil
}

func setUpViper() error {
	viper.SetConfigFile(".env")

	return viper.ReadInConfig()
}
