package config

import (
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	ServerPort int    `mapstructure:"SERVER_PORT"`
	DBURL      string `mapstructure:"DB_URL"`
}

func Load() (*AppConfig, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Try to read .env file, but don't fail if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		// If .env file doesn't exist, that's okay - we'll use environment variables
		if _, ok := err.(*os.PathError); !ok {
			return nil, err
		}
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
