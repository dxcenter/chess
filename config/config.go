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

type UserSourceData map[string]interface{}

func (d UserSourceData) ToInternalUserSourceData() InternalUserSourceData {
	users := Users{}

	usersRaw := d["users"].([]interface{})
	for _, userRawRaw := range usersRaw {
		userRaw := userRawRaw.(map[interface{}]interface{})
		users = append(users, User{
			Login:    userRaw["login"].(string),
			Password: userRaw["password"].(string),
		})
	}
	return InternalUserSourceData{
		Users: users,
	}
}

type InternalUserSourceData struct {
	Users Users
}

func (d UserSourceData) ToDbUserSourceData() (result DbUserSourceData) {
	return DbUserSourceData(d["db_config"].(string))
}

type DbUserSourceData string

type DbCfg struct {
	Driver   string
	Protocol string
	Host     string
	Port     int
	Db       string
	User     string
	Password string
	Path     string
}

type UserSource struct {
	Type string
	Data UserSourceData
}
type UserSources []UserSource

type Dbs map[string]DbCfg

type Config struct {
	Secret      string      `yaml:"secret"`
	Dbs         Dbs         `yaml:"dbs"`
	MyDb        string      `yaml:"my_db"`
	UserSources UserSources `yaml:"user_sources"`
}

var cfg Config

var reloadHooks []func()

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

	for _, reloadHook := range reloadHooks {
		reloadHook()
	}
}

func Get() Config {
	return cfg
}

func AddReloadHook(hook func()) {
	reloadHooks = append(reloadHooks, hook)
}
