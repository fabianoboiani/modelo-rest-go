package configs

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init(env string) {
	var err error

	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../configs/")
	config.AddConfigPath("configs/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on loading configs file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
