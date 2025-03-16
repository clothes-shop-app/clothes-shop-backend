package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort int    `mapstructure:"SERVER_PORT"`
	DBURL      string `mapstructure:"DB_URL"`
}

func Load() (*AppConfig, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
