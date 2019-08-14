package main

import (
	"flag"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"
	"hiris/utils/config"
)

var configFile = flag.String("config", ".conf.yaml", "配置文件路径")

func init() {
	flag.Parse()

}
func main() {
	config.InitConfig(*configFile) //初始化配置
	//initLogrus()//初始化日志
	//initDB()//初始化数据库
}

func initDB() {
	simple.OpenDB(&simple.DBConfiguration{
		Dialect:        "mysql",
		Url:            config.Conf.MySqlUrl,
		MaxIdle:        5,
		MaxActive:      20,
		EnableLogModel: config.Conf.ShowSql,
		Models:         nil, //todo
	})
}

func initLogrus() {
	output, err := simple.NewLogWriter(config.Conf.Logfile)
	if err == nil {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.SetOutput(output)
	} else {
		logrus.Error(err)
	}
}
