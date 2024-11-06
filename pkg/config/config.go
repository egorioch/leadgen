package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func New() *viper.Viper {
	config := viper.New()
	config.SetDefault("configfile", "./config/config.yaml")
	setConfigFile(config)

	return config
}

func setConfigFile(config *viper.Viper) {
	file := config.GetString("configfile")
	fmt.Printf("Set config from file: %s", file)
	config.SetConfigType("yaml")
	config.SetConfigFile(file)
	err := config.ReadInConfig()
	if err != nil {
		log.Printf("Unable to read config from file: %s", err.Error())
		log.Panic("stop")
	}
}
