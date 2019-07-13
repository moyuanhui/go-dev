package tail

import (
	"fmt"
	"go_dev/day11/logagent/config"
	"time"

	"github.com/hpcloud/tail"
)

func InitTail(conf []config.CollectConf) (err error) {
	fileName := "F:\\Go\\src\\my.log"
	tails, err := tail.TailFile(fileName, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file error :", err)
		return
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen filename:", tails.Filename)
			time.Sleep(time.Millisecond)
			continue
		}
		fmt.Println("msg = ", msg)
	}
}
