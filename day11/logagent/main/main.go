package main

import (
	"fmt"
	"go_dev/day11/logagent/config"
	applog "go_dev/day11/logagent/logs"
	"go_dev/day11/logagent/tail"

	"github.com/astaxie/beego/logs"
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
	if err != nil {
		fmt.Printf("load logger failed, err:%v\n", err)
		panic("load logger failed")
		return
	}

	logs.Debug("initialize succ")

	err = tail.InitTail(config.AppConf.CollectConf)
	if err != nil {
		logs.Error("Init tail failed, err:%v\n", err)
		return
	}
}
