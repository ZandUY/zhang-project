package et

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	
)
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	//派一个哨兵一直监视fengyutong这个key的变化
	ch := cli.Watch(context.Background(), "zhangxiangjun")

	//从通道尝试取值
	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v, Key:%v, Value:%v", string(evt.Type), string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}

}
