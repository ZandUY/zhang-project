package taillog

import (
	"fmt"
	"time"

	"code.oldboyedu.com/studygo/logagent/etcd"
)

var tskMgr *tailLogMgr

// tailtask管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf,
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry),
	}

	for _, logEntry := range logEntryConf {
		//初始化的时候，起了多少个tailtask计数
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)

		tskMgr.tskMap[mk] = tailObj

	}
	go tskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理
// 1.配置新增
// 2.配置删除
// 3.配置变更
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println("新的配置来了!\n", newConf)

			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有不需要操作
					continue
				} else { //新增的

					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}

			for k, _ := range t.tskMap {
				isDelete := true
				for _, c2 := range newConf {
					if string(c2.Path)+"_"+string(c2.Topic) == k {
						isDelete = false
						continue
					}
				}
				if isDelete {
					//把c1对应的这个tailObj给停掉
					mk := fmt.Sprintf("%s", k)
					fmt.Println(mk)
					t.tskMap[mk].cancelFunc()
					delete(t.tskMap, mk)

				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

// 向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
