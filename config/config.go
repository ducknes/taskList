package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() *Settings {
	env, err := readEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Env: %s\n", env)
	var settings Settings
	configFilePath := configPath(env)
	configFileBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = json.Unmarshal(configFileBytes, &settings); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return &settings
}

func configPath(env Env) string {
	return fmt.Sprintf("./.config/%s.json", env)
}

func readEnv() (Env, error) {
	env := os.Getenv(EnvVarName)
	switch Env(env) {
	case Env(""):
		return Local, nil
	case Production:
		return Production, nil
	default:
		return "", fmt.Errorf("unknown environmwnt '%s'", os.Args[1])
	}
}
