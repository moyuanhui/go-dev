package logs

import (
	"encoding/json"
	"fmt"
	appconfig "go_dev/day11/logagent/config"

	"github.com/astaxie/beego/logs"
)

func converLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

func InitLogger() (err error) {

	config := make(map[string]interface{})
	config["filename"] = appconfig.AppConf.LogPath
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)

	if err != nil {
		fmt.Println("marshal failed ,err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
