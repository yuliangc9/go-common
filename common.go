package common

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
)

func TestPrint() {
	fmt.Println("vim-go")
	beego.Info("hehe")
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Info("hehe")
}

var Config config.Configer

func initLogWithConfig() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(4)

	level := Config.DefaultInt("log::level", 7)

	if Config.DefaultBool("log::console", true) {
		configStr := fmt.Sprintf(`{"level":%d,"color":true}`, level)
		logs.SetLogger(logs.AdapterConsole, configStr)
	}

	if file := Config.String("log::file"); file != "" {
		configStr := fmt.Sprintf(
			`{"level":%d,"filename":"%s","hourly":true, "maxhours":%d,"daily":false,"maxlines":%d,"maxsize":%d}`,
			level, file, Config.DefaultInt("log::max_hours", 7*24), 1<<31, 1<<31)
		logs.SetLogger(logs.AdapterFile, configStr)
	}
}

func initLogDefault() {
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(4)

	logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
}

func init() {
	configFile := "./app.conf"
	if f := os.Getenv("CONF_FILE"); f != "" {
		configFile = f
	}

	var err error
	Config, err = config.NewConfig("ini", configFile)
	if err != nil {
		logs.Warn("open config file fail", err.Error())

		initLogDefault()
	} else {
		initLogWithConfig()
	}
}

// Emergency logs a message at emergency level.
func Emergency(f interface{}, v ...interface{}) {
	logs.Emergency(f, v...)
}

// Alert logs a message at alert level.
func Alert(f interface{}, v ...interface{}) {
	logs.Alert(f, v...)
}

// Error logs a message at error level.
func Error(f interface{}, v ...interface{}) {
	logs.Error(f, v...)
}

// Warn compatibility alias for Warning()
func Warn(f interface{}, v ...interface{}) {
	logs.Warn(f, v...)
}

// Notice logs a message at notice level.
func Notice(f interface{}, v ...interface{}) {
	logs.Notice(f, v...)
}

func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v...)
}

func Debug(f interface{}, v ...interface{}) {
	logs.Debug(f, v...)
}
