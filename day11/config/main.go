package main

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

func main() {
	config, err := config.NewConfig("ini", "F:\\Go\\src\\logcollect.conf")
	if err != nil {
		fmt.Println("new config failed err:", err)
		return
	}

	port, err := config.Int("server::port")
	if err != nil {
		fmt.Println("read server::port failed,err", err)
		return
	}

	fmt.Println("Port:", port)
	logLevel := config.String("log::log_level")
	// if err != nil {
	// 	fmt.Println("read log_level failed,err:", err)
	// 	return
	// }
	fmt.Println("log log_level:", logLevel)
	logPath := config.String("log::log_path")
	fmt.Println("log_path:", logPath)
}
