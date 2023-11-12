package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Settings struct {
	Port               string `json:"port"`
	PostgresConnection string `json:"postgresConnection"`
}

func Init() {
	viper.AddConfigPath(".config/")
	viper.SetConfigType("json")
	viper.SetConfigName(configFile())
	err := viper.WriteConfig()
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
