package config

import (
	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

func InitiliazeConfig() {

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}

}
