package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"time"
)

func main() {

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 禁用证书验证
			},
		},
	}

	// 创建一个上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 创建 Elasticsearch 客户端
	client, err := elastic.NewClient(
		elastic.SetURL("https://172.21.24.15:9200"), // 替换为你的 Elasticsearch 地址
		elastic.SetSniff(false),                     // 禁用嗅探
		//elastic.SetHealthcheck(false),               // 禁用健康检查
		elastic.SetBasicAuth("elastic", "6GS6&^pR5Pzh20N7"),
		elastic.SetHttpClient(httpClient),
	)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// 检查连接
	_, _, err = client.Ping("https://172.21.24.15:9200").Do(ctx)
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch: %s", err)
	}

	fmt.Println("Elasticsearch is up and running!")
}
