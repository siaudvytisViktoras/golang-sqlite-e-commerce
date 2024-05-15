package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseFilePath string
	ServerPort       string
}

var AppConfig Config

func LoadConfig() {
	viper.SetDefault("~/Go/golang-sqlite-e-commerce/config", "fist.db")
	viper.SetDefault("ServerPort", "3000")

	viper.SetConfigName("first") // Configuration file name (without extension)
	viper.AddConfigPath(".")     // Search for the config file in the current directory

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		panic(err)
	}
}
