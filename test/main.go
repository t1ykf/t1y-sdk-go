package main

import (
	"log"

	"github.com/t1ykf/t1y-sdk-go"
)

// 表结构
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func main() {
	// 初始化 SDK 配置
	client := t1y.Init("https://example.com", 1001, "api_key", "secret_key")

	// 添加一条数据
	response, err := client.CreateOne("student", &Student{Name: "王华", Age: 21, Sex: "男"})
	if err != nil {
		log.Printf("fail: %v", err)
		return
	}
	log.Printf("success: %v", response)
}
