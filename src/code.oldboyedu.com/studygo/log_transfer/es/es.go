package es

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

var (
	client *elasticsearch.Client
	ch     = make(chan *LogData, 100000)
)

// 初始化es，准备接收kafka发来的数据
func Init(address string) (err error) {
	// 创建 Elasticsearch 客户端
	if !strings.HasPrefix(address, "http://") {
		address = "https://" + address
	}

	client, err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{address},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 禁用证书验证
		},
		Username: "elastic",              // 替换为你的用户名
		Password: "+VKuz*e2+G0g+Rip0YWD", // 替换为你的密码
	})
	if err != nil {
		log.Fatalf("Error creating the client: %v", err)
		return
	}
go sendToES()
	//发送数据到ES
	return nil
}

type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func SendToESChan(msg *LogData) {
	ch <- msg
}
func sendToES() {
	for {
		select {
		case msg :=<- ch:
			fmt.Println(msg.Topic)
			dataTmp, err := json.Marshal(msg)
			// dataTmp := []byte(`{"name":"zhangxiangjun","age":21,"married":true}`)
			// p1 := student{
			// 	Name:    "fengyutong",
			// 	Age:     19,
			// 	Married: false,
			// }
			// 将 student 对象转换为 JSON
			// dataTmp, err := json.Marshal(p1)
			// var err error
			// if err != nil {
			// 	log.Fatalf("Error marshaling JSON: %v", err)
			// }
			// fmt.Printf("Type :%T Value:%v\n", dataTmp, string(dataTmp))
			// 执行索引请求
			req := esapi.IndexRequest{
				Index:      msg.Topic,                 // 索引名称
				DocumentID: "",                       // 可选：设置文档 ID
				Body:       bytes.NewReader(dataTmp), // 请求体
				Refresh:    "true",                   // 可选：刷新索引
			}

			// 执行索引操作
			res, err := req.Do(context.Background(), client)
			if err != nil {
				log.Fatalf("Error indexing document: %v\n", err)
			}
			defer res.Body.Close()

			// 打印响应
			if res.IsError() {
				log.Printf("[%s] Error indexing document", res.Status())
			} else {
				// fmt.Printf("index %s 已经发送到ES\n", req.Index)
				fmt.Println()
			}
		default:
			time.Sleep(time.Second)
		}
	}

}
