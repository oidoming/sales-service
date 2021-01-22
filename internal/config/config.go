package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configurations struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func Load() *Configurations {
	var config Configurations

	viper.SetConfigName("config")
	viper.AddConfigPath("../.")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return &config
}
