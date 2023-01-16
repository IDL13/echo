package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	username string `yaml:"username"`
	password string `yaml:"password"`
	host     string `yaml:"host"`
	port     string `yaml:"port"`
	database string `yaml:"database"`
}

func ReadConf() *Config {
	con := &Config{}

	conf_date, err := os.ReadFile("conf.yml")
	if err != nil {
		log.Fatal("Read conf error")
	}

	yaml.Unmarshal(conf_date, con)

	return con
}
