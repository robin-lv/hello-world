package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
)

const ELASTICSEARCH_URL = "ELASTICSEARCH_URL"

var es *elasticsearch.Client

func InitEnv() {
	os.Setenv(ELASTICSEARCH_URL, "http://localhost:9200")
	// 创建 Elasticsearch 客户端
	var err error
	es, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
}

// 定义样本数据结构
type Person struct {
	Name string `json:"name"`
	UID  int    `json:"uid"`
}

// 定义索引结构
type IndexMeta struct {
	Index struct {
		ID int `json:"_id"`
	} `json:"index"`
}

func main() {
	InitEnv()

	// 打开 JSON 文件
	file, err := os.Open("./sample_data.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var batch []any
	buf := &bytes.Buffer{}

	for scanner.Scan() {
		// 读取 index 元数据行
		var indexMeta IndexMeta
		if err := json.Unmarshal(scanner.Bytes(), &indexMeta); err != nil {
			log.Fatalf("Failed to unmarshal index metadata: %v", err)
		}
		buf.Write(scanner.Bytes())
		buf.WriteByte('\n')
		// 读取数据行
		if !scanner.Scan() {
			log.Fatalf("Expected data line after index metadata")
		}
		var person Person
		if err := json.Unmarshal(scanner.Bytes(), &person); err != nil {
			log.Fatalf("Failed to unmarshal person data: %v", err)
		}
		buf.Write(scanner.Bytes())
		buf.WriteByte('\n')
		// 将数据添加到批次
		batch = append(batch, &indexMeta, &person)
		if len(batch) >= 20000 {
			processBatch(buf)
			batch = nil
			buf.Reset()
		}
	}

}

func processBatch(buf *bytes.Buffer) {
	// 索引文档
	res, err := es.Bulk(
		buf,

		es.Bulk.WithIndex("gz_mb-username"),
	)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()
	fmt.Println("Search results:", res.Status(), string(buf.Bytes()[0:20]))
}
