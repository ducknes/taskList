package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Settings struct {
	Port               string `json:"port"`
	PostgresConnection string `json:"postgresConnection"`
}

func Init() {
	viper.AddConfigPath(".config/")
	viper.SetConfigType("json")
	viper.SetConfigName(configFile())
	err := viper.ReadInConfig()
	viper.WatchConfig()
	if err != nil {
		panic(fmt.Sprintf("Couldn't write config. error: %v", err))
	}
}

func configFile() string {
	env := os.Getenv(EnvVarName)
	if env == "" {
		return string(Local)
	}

	return env
}
