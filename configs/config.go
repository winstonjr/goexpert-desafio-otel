package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io/fs"
)

type conf struct {
	WeatherApiKey  string `mapstructure:"WEATHER_API_KEY"`
	InternalApiURI string `mapstructure:"INTERNAL_API_URI"`
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")

	// Environment variables will override the .env values
	viper.AutomaticEnv()

	// Read the .env file
	if err := viper.ReadInConfig(); err != nil {
		// If the .env file is not found, continue without error
		fmt.Printf("Loading config from .env file %T\n", err)
		var pathError *fs.PathError
		if !errors.As(err, &pathError) {
			return nil, err
		}
	}

	var cfg *conf
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
