package main

import (
	"code.oldboyedu.com/studygo/day06/mylogger"
)
var log mylogger.Logger
func main() {
	log = mylogger.NewConsoleLogger("DEBUG")
	log = mylogger.NewFileLogger("INFO", "./", "Feng.log", 10*1024*1024)
	for {
		id := 10010
		name := "理想"
		log.Debug("这是一条DEBUG日志,id:%d,name:%s", id, name)
		log.Info("这是一条Info日志")
		log.Warning("这是一条Waring日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
	}

}
