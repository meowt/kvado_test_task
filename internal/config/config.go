package config

import (
	"log"

	"github.com/spf13/viper"
)

// MustLoad reads config file
func MustLoad() {
	//Specifying config file
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	//Reading config
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	log.Println("Config loaded")
}
