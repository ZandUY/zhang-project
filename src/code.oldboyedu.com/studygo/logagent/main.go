package main

import (
	"fmt"
	"time"

	"code.oldboyedu.com/studygo/logagent/conf"
	"code.oldboyedu.com/studygo/logagent/kafka"
	taillog "code.oldboyedu.com/studygo/logagent/tail_log"
	"gopkg.in/ini.v1"
)

var (
	// cfg, err = ini.Load("./conf/config.ini")
	cfg = new(conf.AppConf)
)

func run() {
	for {
		select {
		case line := <-taillog.ReadChan():
			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
			fmt.Printf("%v\n", line.Text)
		default:
			time.Sleep(time.Second)
		}

	}
}
func main() {
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("ini file open filed,err:%v", err)
		return
	}
	//1.初始化kafak连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init kafak failed,err:%v\n", err)
		return
	}
	//2.打开日志文件准备收集
	err = taillog.Init(cfg.TaillogConf.FileName)
	if err != nil {
		fmt.Printf("init taillog failed,err:%v\n", err)
	}
	fmt.Println("tail连接成功")
	run()
}
