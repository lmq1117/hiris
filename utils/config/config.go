package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *Config

type Config struct {
	Env      string `yaml:"Env"`      //环境：prod env
	MySqlUrl string `yaml:"MySqlUrl"` //数据库连接地址
	ShowSql  bool   `yaml:"ShowSql"`  //是否显示Sql
	Logfile  string `yaml:"Logfile"`  //日志文件
}

func InitConfig(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
		return
	}
	Conf = &Config{}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		logrus.Error(err)
	}
}
