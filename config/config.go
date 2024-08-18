package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Version string `mapstructure:"VERSION"`
	DB      struct {
		Host     string `mapstructure:"DB_HOST"`
		Port     int    `mapstructure:"DB_PORT"`
		Username string `mapstructure:"DB_USERNAME"`
		Password string `mapstructure:"DB_PASSWORD"`
		Name     string `mapstructure:"DB_NAME"`
	}
	BinanceAPIKey    string `mapstructure:"BINANCE_API_KEY"`
	BinanceSecretKey string `mapstructure:"BINANCE_SECRET_KEY"`
	DatabaseDSN      string
}

var Settings Config

func init() {
	viper.SetConfigFile(".env")
	// viper.SetConfigType("env") // Required if there is not extension in the name.

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file")
	}

	if err := viper.Unmarshal(&Settings); err != nil {
		panic("Failed to unmarshal the enviroments variable to Settings struct")
	}

	Settings.DB.Host = viper.GetString("DB_HOST")
	Settings.DB.Port = viper.GetInt("DB_PORT")
	Settings.DB.Username = viper.GetString("DB_USERNAME")
	Settings.DB.Password = viper.GetString("DB_PASSWORD")
	Settings.DB.Name = viper.GetString("DB_NAME")

	Settings.DatabaseDSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		Settings.DB.Host, Settings.DB.Username, Settings.DB.Password,
		Settings.DB.Name, Settings.DB.Port)
}
