package main

import (
	"fmt"
	"go_dev/day11/logagent/config"
	applog "go_dev/day11/logagent/logs"
)

func main() {

	//加载配置
	fileName := "./config/logagent.conf"
	confType := "ini"
	err := config.LoadConfig(confType, fileName)
	if err != nil {
		fmt.Println("load failed config,err:", err)
		return
	}

	//初始化日志
	err = applog.InitLogger()

}
