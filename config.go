package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ReadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	if os.Getenv("CONFIG_PATH") != "" {
		viper.AddConfigPath(os.Getenv("TELEGRAM_ENV_PATH"))
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
