package utils

import "github.com/spf13/viper"

type Config struct {
	UserServiceServerHost string `mapstructure:"UserServiceServerHost"`
	UserServiceServerPort int    `mapstructure:"UserServiceServerPort"`
}

var Configuration Config

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile("./conf/config.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
