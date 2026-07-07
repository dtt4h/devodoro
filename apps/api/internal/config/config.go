package config

type Config struct {
	Env      string   `yaml:"env"`
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
	Logger   Logger   `yaml:"logger"`
}

type Database struct {
	URL string `yaml:"url"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Logger struct {
	Level string `yaml:"level"`
}
