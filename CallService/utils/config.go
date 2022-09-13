package utils

import "github.com/spf13/viper"

type Config struct {
	UserServiceServerHost string `mapstructure:"UserServiceServerHost"`
	UserServiceServerPort int    `mapstructure:"UserServiceServerPort"`
	MainServiceServerHost string `mapstructure:"MainServiceServerHost"`
	MainServiceServerPort int    `mapstructure:"MainServiceServerPort"`
}

var Configuration Config

func LoadConfig() (err error) {
	viper.SetConfigFile("./conf/config.env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&Configuration)
	return
}
