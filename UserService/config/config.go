package config

import (
	"sync"
	"time"
)

var (
	once      sync.Once
	singleton *Config
)

func SetConfig(c *Config) *Config {
	once.Do(func() {
		singleton = c
	})

	return singleton
}

func GetConfig() *Config {
	if singleton == nil {
		panic("config not set")
	}

	return singleton
}

type Config struct {
	Env             string
	MysqlCfg        Mysql
	ServerCfg       Server
	JWTSecret       string
	EnableProfiling bool
}

type Mysql struct {
	Host            string
	Port            string
	Username        string
	Password        string
	Name            string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int // time in minute
	DB_SSLMode      string
	IsEnabledLog    bool
}

type Server struct {
	Port    string
	Timeout time.Duration
}
