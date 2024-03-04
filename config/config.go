package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GRPCPORT string `mapstructure:"GRPC_CLIENT_PORT"`
}

func Init() *Config {
	viper.SetConfigFile("config")
	viper.AddConfigPath("config")
	viper.SetConfigFile("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

	config := Config{}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Panicln("Error marshaling config file")
	}
	return &config
}
