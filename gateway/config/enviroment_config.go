package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	path       = "."
	configName = "app"
	configType = "env"
)

type Configuration struct {
	Port             string `mapstructure:"PORT"`
	AccountClientUrl string `mapstructure:"ACCOUNT_CLIENT_URL"`
}

func NewConfig() Configuration {

	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	conf := &Configuration{}

	err = viper.Unmarshal(conf)

	if err != nil {
		fmt.Printf("unable decode into config struct, %v", err)
	}
	return *conf
}
