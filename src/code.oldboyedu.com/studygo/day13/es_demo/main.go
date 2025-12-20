package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/esapi"
)

type student struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func main() {
	// 创建 Elasticsearch 客户端
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"https://127.0.0.1:9200"},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 禁用证书验证
		},
		Username: "elastic",              // 替换为你的用户名
		Password: "+VKuz*e2+G0g+Rip0YWD", // 替换为你的密码
	})
	if err != nil {
		log.Fatalf("Error creating the client: %v", err)
	}
	fmt.Println("Connected to Elasticsearch successfully.")

	// 创建一个 student 对象
	p1 := student{
		Name:    "fengyutong",
		Age:     19,
		Married: false,
	}
	// 将 student 对象转换为 JSON
	data, err := json.Marshal(p1)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// 执行索引请求
	req := esapi.IndexRequest{
		Index:      "user",                // 索引名称
		DocumentID: "",                    // 可选：设置文档 ID
		Body:       bytes.NewReader(data), // 请求体
		Refresh:    "true",                // 可选：刷新索引
	}

	// 执行索引操作
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatalf("Error indexing document: %v", err)
	}
	defer res.Body.Close()

	// 打印响应
	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	} else {
		fmt.Printf("Indexed user to index %s\n", req.Index)
	}
}
