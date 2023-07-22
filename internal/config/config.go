package config

import (
	"os"

	"github.com/IDL13/echo/pkg/utils"
	"gopkg.in/yaml.v3"
)

func GetConf() *Config {
	c := &Config{}
	info, err := os.ReadFile("./../conf.yaml")
	if err != nil {
		utils.Loger(err)
	}
	yaml.Unmarshal(info, c)
	return c
}
