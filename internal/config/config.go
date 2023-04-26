package config

import (
	"os"

	"github.com/IDL13/echo/pkg/utils"
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
	info, err := os.ReadFile("conf.yaml")
	if err != nil {
		utils.Loger(err)
	}
	yaml.Unmarshal(info, c)
	return c
}
