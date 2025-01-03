package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"os"
	"time"
)

func genName() string {
	// 初始化 gofakeit
	gofakeit.Seed(0)
	if rand.Float32() < 0.7 {
		return gofakeit.Name()
	}

	return generateRandomName()

}

// 定义样本数据结构
type SampleData struct {
	Name string `json:"name"`
	UID  int    `json:"uid"`
}

// 定义索引结构
type Index struct {
	Index struct {
		ID int `json:"_id"`
	} `json:"index"`
}

// 随机生成中英文名字
func generateRandomName() string {

	chineseNames := []string{"小明", "小红", "大卫", "汉堡", "小丽", "小刚", "小华", "小芳", "小杰", "小强"}

	// 合并名字列表

	return chineseNames[rand.Intn(len(chineseNames))]
}

// 生成样本数据
func generateSampleData(numSamples int) (data []interface{}, err error) {
	for uid := 1; uid <= numSamples; uid++ {

		index := Index{}
		index.Index.ID = uid

		sample := SampleData{Name: genName(), UID: uid}

		data = append(data, index)
		data = append(data, sample)
	}
	return
}

// 写入文件
func writeToFile(data []interface{}, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, entry := range data {
		if err := json.NewEncoder(file).Encode(entry); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	numSamples := 1000000
	sampleData, err := generateSampleData(numSamples)
	if err != nil {
		fmt.Println("生成样本数据时出错:", err)
		return
	}

	outputFilename := "./sample_data.json"
	if err = writeToFile(sampleData, outputFilename); err != nil {
		fmt.Println("写入文件时出错:", err)
		return
	}

	fmt.Printf("生成了 %d 条样本数据，已保存到 %s\n", numSamples, outputFilename)
}
