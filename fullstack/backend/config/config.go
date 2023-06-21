package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port              string
	JWTSecret         string
	DatabaseURL       string
	DatabaseName      string
	AnalytixModelPath string
}

func LoadConfig() (*Config, error) {
	// Set default values for the configuration options
	viper.SetDefault("port", "8080")
	viper.SetDefault("jwt_secret", "secret")
	viper.SetDefault("database_url", "mongodb://localhost:27017")
	viper.SetDefault("database_name", "auth_app")
	viper.SetDefault("analytix_model_path", "analytix/model")

	// Set the name of the configuration file
	viper.SetConfigName("config")

	// Set the paths to search for the configuration file
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/auth_app")
	viper.AddConfigPath("/etc/auth_app/")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No configuration file found, using defaults.")
		} else {
			return nil, err
		}
	}

	// Create the config object and set its values from the configuration file or the defaults
	config := &Config{
		Port:              viper.GetString("port"),
		JWTSecret:         viper.GetString("jwt_secret"),
		DatabaseURL:       viper.GetString("database_url"),
		DatabaseName:      viper.GetString("database_name"),
		AnalytixModelPath: viper.GetString("analytix_model_path"),
	}

	return config, nil
}
