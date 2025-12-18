package main

import (
	"fmt"
	"sync"
	"time"

	"code.oldboyedu.com/studygo/logagent/conf"
	"code.oldboyedu.com/studygo/logagent/etcd"
	"code.oldboyedu.com/studygo/logagent/kafka"
	taillog "code.oldboyedu.com/studygo/logagent/tail_log"
	"gopkg.in/ini.v1"
)

var (
	// cfg, err = ini.Load("./conf/config.ini")
	cfg = new(conf.AppConf)
)

// func run() {
// 	for {
// 		select {
// 		case line := <-taillog.ReadChan():
// 			kafka.SendToKafka(cfg.KafkaConf.Topic, line.Text)
// 			fmt.Printf("%v\n", line.Text)
// 		default:
// 			time.Sleep(time.Second)
// 		}

//		}
//	}

func main() {
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("ini file open filed,err:%v", err)
		return
	}
	//1.初始化kafak连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.MaxSize)
	if err != nil {
		fmt.Printf("init kafak failed,err:%v\n", err)
		return
	}
	// //2.打开日志文件准备收集
	// err = taillog.Init(cfg.TaillogConf.FileName)
	// if err != nil {
	// 	fmt.Printf("init taillog failed,err:%v\n", err)
	// }
	// fmt.Println("tail连接成功")
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.TimeOut)*time.Second)
	if err != nil {
		fmt.Printf("Etcd init failed,err :%v\n", err)
		return
	}

	//从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	if err != nil {
		fmt.Printf("Etcd getConf failed,err :%v\n", err)
		return
	}

	for _, v := range logEntryConf {
		fmt.Println(v.Path, v.Topic)
	}

	//收集日志发往Kafka
	taillog.Init(logEntryConf)
	//派一个哨兵
	var wg sync.WaitGroup
	wg.Add(1)
	newConfChan := taillog.NewConfChan()
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan)
	
	wg.Wait()

}
