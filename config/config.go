package config

import (
	"github.com/spf13/viper"
)

func Load() {

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

}
