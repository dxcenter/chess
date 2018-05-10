package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type User struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
}

type Users []User

type Config struct {
	Secret string `yaml:"secret"`
	Users  Users  `yaml:"users"`
}

var cfg Config

func checkErr(err error) {
	if err == nil {
		return
	}

	panic(err)
}

func Reload() {
	configData, err := ioutil.ReadFile("config.yaml")
	checkErr(err)

	err = yaml.Unmarshal([]byte(configData), &cfg)
	checkErr(err)
}

func Get() Config {
	return cfg
}
