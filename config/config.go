package config

import (
	"fmt"
	"go_scaffold_project/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Port      string
	JwtSecret string
)

func logsinit() {
	logLevel := viper.GetString("logs.level")
	// 设置日志级别
	switch logLevel {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
	}
	logs.Info(nil, fmt.Sprint("Set log level to : ", logrus.GetLevel()))
	// 日志中显示文件名
	SetReportCaller := viper.GetBool("logs.showcaller")
	logrus.SetReportCaller(SetReportCaller)
	// 时间显示格式
	TimestampFormat := viper.GetString("logs.timeformat")
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimestampFormat})

}

func configload() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logs.Error(nil, "Filed to read config")
	}
}

func gininit() {
	///todo
	Port = viper.GetString("program.port")
}

func jwtinit() {
	JwtSecret = viper.GetString("program.jwtsecret")
}

func init() {
	// 加载程序配置
	logs.Debug(nil, "Config init...")
	configload()
	logsinit()
	gininit()
	jwtinit()
	logs.Debug(nil, "Config init completed")
}
