package config

type MainConfig struct {
	App AppConfig
}

type AppConfig struct {
	Host string
	Port int
}
