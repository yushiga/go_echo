package config

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config Conf

type Environment struct {
	Development Conf `yaml:"development"`
	Staging     Conf `yaml:"staging"`
	Production  Conf `yaml:"production"`
}

type Conf struct {
	Database Database `yaml:"database"`
	Logger   Logger   `yaml:"log"`
}

/*
DB設定
*/
type Database struct {
	Driver   string `yaml:"driver"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

/*
ログ設定
*/
type Logger struct {
	File      string `yaml:"file"`
	MaxAge    int    `yaml:"maxAge"`
	MaxBackup int    `yaml:"maxBackup"`
}

/*
環境設定
*/
func SetEnvironment(env string) {
	buf, err := ioutil.ReadFile("config/application.yml")
	if err != nil {
		panic(err)
	}

	var environment Environment

	err = yaml.Unmarshal(buf, &environment)
	if err != nil {
		panic(err)
	}

	if env == "production" {
		Config = environment.Production
		fmt.Println("===== Start " + env + " ======")
	} else if env == "staging" {
		fmt.Println("===== Start " + env + " ======")
		Config = environment.Staging
	} else {
		fmt.Println("===== Start " + env + " ======")
		Config = environment.Development
	}
}
