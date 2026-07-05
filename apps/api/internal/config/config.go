package config

type Config struct {
	Env      string   `yaml:"env"`
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

type Database struct {
	URL  string `yaml:"url"`
	Name string `yaml:"name"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
