package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	UserServiceServerHost string `mapstructure:"UserServiceServerHost"`
	UserServiceServerPort int    `mapstructure:"UserServiceServerPort"`
	MainServiceServerHost string `mapstructure:"MainServiceServerHost"`
	MainServiceServerPort int    `mapstructure:"MainServiceServerPort"`
	HMAC_KEY              string `mapstructure:"HMAC_KEY"`
}

func (c Config) getMainServiceServerHost() (string, error) {
	return c.MainServiceServerHost, nil
}

func (c Config) getMainServiceServerPort() (int, error) {
	return c.MainServiceServerPort, nil
}

func (c Config) getUserServiceServerHost() (string, error) {
	return c.UserServiceServerHost, nil
}

func (c Config) getUserServiceServerPort() (int, error) {
	return c.UserServiceServerPort, nil
}

func (c Config) getHMACKey() (string, error) {
	return c.HMAC_KEY, nil
}

var Configuration = Config{}

func init() {
	err := LoadConfig()
	if err != nil {

	}
}

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
