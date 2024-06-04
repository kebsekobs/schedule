package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	DBConfig *viper.Viper
)

func init() {
	fmt.Printf("Loading configuration logics...\n")
	DBConfig = initConfig("db")
}

func initConfig(name string) *viper.Viper {
	config := viper.New()
	config.SetConfigName(name)
	config.AddConfigPath("./internal/config")
	config.SetConfigType("yaml")
	err := config.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to get the configuration.")
	}
	return config
}
