package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

type LogEntry struct {
	Path  string `json:"path"`  //日志存放的路径
	Topic string `json:"topic"` //日志发往Kafka中哪个Topic
}

func Init(addr string, timeOut time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr},
		DialTimeout: timeOut,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Printf("连接Etcd成功\n")
	return
}
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// var logEntryConf=make([]*LogEntry,5,10)
	for _, ev := range resp.Kvs {
		// fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unMarshal etcd  value failed,err:%v\n", err)
			return
		}
	}
	return
}
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	ch := cli.Watch(context.Background(), key)
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%s, Key:%v, Value:%v\n", string(evt.Type), string(evt.Kv.Key), string(evt.Kv.Value))
			//通知taillog.tskMgr
			//1.判断操作的类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				//如果是删除操作，手动传递一个空的配置项
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("json unmarshal failed,err :%v", err)
					continue
				}

			}
			fmt.Printf("get new conf:%v\n", newConf)
			newConfCh <- newConf
		}

	}
}
