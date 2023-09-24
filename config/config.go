package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	ConfigFile = "config/config.yml"
	configType = "yml"
)

type (
	Config struct {
		Server   Server   `mapstructure:"server"`
		Database Database `mapstructure:"database"`
	}

	Server struct {
		Host     string `mapstructure:"host"`
		Env      string `mapstructure:"env"`
		UseRedis bool   `mapstructure:"useRedis"`
		Port     int    `mapstructure:"port"`
	}

	Database struct {
		Driver   string `mapstructure:"driver"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	}
)

func NewConfig() *Config {
	initConfig()
	conf := &Config{}
	err := viper.Unmarshal(conf)
	if err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	return conf
}

func initConfig() {
	var configFile string
	configFile = ConfigFile

	viper.SetConfigType(configType)
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}
}
