package main

import (
	"fmt"

	"code.oldboyedu.com/studygo/log_transfer/conf"
	"code.oldboyedu.com/studygo/log_transfer/es"
	"code.oldboyedu.com/studygo/log_transfer/kafka"

	"gopkg.in/ini.v1"
)

func main() {

	//0.加载配置文件
	var cfg conf.LogTransferCfg
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("配置文件加载失败,err:%v!\n", err)
		return
	}
	fmt.Printf("配置文件加载成功!   %#v\n", cfg)
	//初始化ES
	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Printf("初始化ES失败,err:%v\n", err)
		return
	}
	fmt.Println("初始化ES成功")
	//初始化kafka
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("初始化kafka失败,err:%v\n", err)
		return
	}
	select {}
	//2.从kafka取日志数据
	//3.发往ES
}
