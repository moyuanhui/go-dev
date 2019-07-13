package config

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego/config"
)

var (
	AppConf *Config
)

func LoadConfig(confType, fileName string) (err error) {

	config, err := config.NewConfig(confType, fileName)
	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}

	AppConf = &Config{}
	AppConf.LogLevel = config.String("logs::log_level")
	if len(AppConf.LogLevel) == 0 {
		AppConf.LogLevel = "debug"
	}

	AppConf.LogPath = config.String("logs::log_path")
	if len(AppConf.LogPath) == 0 {
		AppConf.LogPath = "./logs"
	}

	err = loadCollectConf(config)
	if err != nil {
		fmt.Println("load collect conf failed, err:", err)
		return
	}

	return
}

func loadCollectConf(conf config.Configer) (err error) {

	var cc CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invlid collect::log_path")
		return
	}
	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invlid collect::topic")
		return
	}

	AppConf.CollectConf = append(AppConf.CollectConf, cc)

	return
}
