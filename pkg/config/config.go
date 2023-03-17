package config

import "github.com/spf13/viper"

type Config struct {
	TelegramBotToken string
	DatabaseURL      string
}

func Init() (*Config, error) {
	if err := setUpViper(); err != nil {
		return nil, err
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

	if err := viper.BindEnv("DATABASE_URL"); err != nil {
		return err
	}
	cfg.DatabaseURL = viper.GetString("DATABASE_URL")

	return nil
}

func setUpViper() error {
	viper.SetConfigFile(".env")

	return viper.ReadInConfig()
}
