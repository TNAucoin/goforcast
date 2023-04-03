package config

import (
	"github.com/spf13/viper"
	"log"
)

func ReadEnvValue(key string) string {
	viper.SetConfigFile("./.env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
