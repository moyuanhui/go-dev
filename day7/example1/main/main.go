package main

import (
	"fmt"
	"go_dev/day7/example1/balance"
	"math/rand"
	"os"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 3; i++ {
		host := fmt.Sprintf("192.168.100.%d", rand.Intn(255))
		port := 8080
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	// conf := "random"
	// var balancer balance.IBalancer
	// if len(os.Args) > 1 {
	// 	conf = os.Args[1]
	// }

	// if conf == "random" {
	// 	balancer = &balance.RandomBalance{}
	// 	fmt.Println("use random balance")
	// } else if conf == "runxun" {
	// 	balancer = &balance.RunxunBalance{}
	// 	fmt.Println("use runxun balance")

	// }

	conf := "random"
	// var balancer balance.IBalancer
	if len(os.Args) > 1 {
		conf = os.Args[1]
	}

	for {
		// inst, err := balancer.DoBalance()
		inst, err := balance.DoBalance(conf, insts)
		if err != nil {
			fmt.Println("do balance err:", err)
			continue
		}
		// fmt.Printf("%s:%d", inst.GetHost(), inst.GetPort())
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
