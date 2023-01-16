package config

import (
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

var instance *Config
var once sync.Once

func GetConf() *Config {
	once.Do(func() {
		conf_date, err := os.ReadFile("conf.yml")
		if err != nil {
			log.Fatal("Read conf error")
		}

		yaml.Unmarshal(conf_date, instance)
	})
	return instance
}
