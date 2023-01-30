package Config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

func GetConf() *Config {
	c := &Config{}
	info, err := os.ReadFile("Z:/PROGRAMS/GO/echo/conf.yaml")
	if err != nil {
		log.Fatal("Error in read conf file")
	}
	yaml.Unmarshal(info, c)
	return c
}
