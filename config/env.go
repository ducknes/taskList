package config

type Env string

const EnvVarName = "ENV"

const (
	Local      = Env("local")
	Production = Env("production")
)
