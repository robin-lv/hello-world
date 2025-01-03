package fsearch

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
	"strings"
	"testing"
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

func BenchmarkSearch(b *testing.B) {
	InitEnv()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			benchSearch(b)
		}
	})
}
func benchSearch(b *testing.B) {
	// 搜索文档
	searchRes, err := es.Search(
		es.Search.WithIndex("gz_mb-username"),
		es.Search.WithBody(strings.NewReader(`{"query": {"match": {"name": "小"}}}`)),
		es.Search.WithSize(25),
		es.Search.WithPretty(),
		es.Search.WithTimeout(100),
	)
	if err != nil {
		b.Errorf("Unexpected error when getting a response: %s", err)
	}
	searchRes.Body.Close()
}
func TestExample(t *testing.T) {
	InitEnv()

	// 索引文档
	res, err := es.Index(
		"index_v1",
		strings.NewReader(`{
			"title": "Test Document"
		}`),
		es.Index.WithDocumentID("123"),
	)
	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
	}
	defer res.Body.Close()

	// 搜索文档
	searchRes, err := es.Search(
		es.Search.WithIndex("gz_mb-username"),
		es.Search.WithBody(strings.NewReader(`{"query": {"match": {"title": "小"}}}`)),
	)
	if err != nil {
		log.Fatalf("Error searching documents: %s", err)
	}
	defer searchRes.Body.Close()

	fmt.Println("Search results:", searchRes)
}
