package main

import (
	"github.com/spf13/viper"
)

func initViper() {
	viper.SetConfigName("server-config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Config file found and successfully parsed
}
