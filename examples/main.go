package main

import (
	"encoding/json"
	"log"

	"github.com/t1ykf/t1y-sdk-go"
)

// 表结构
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

// 响应
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	// 初始化 SDK 配置
	client := t1y.Init("http://dev.t1y.net/api", 1001, "2c6118c4e02b40fe96f5c40ee1dc5561", "650bd657da0243b282d9cab6d75a80ff")

	// 创建一条数据
	response, err := client.CreateOne("student", &Student{Name: "王华", Age: 21, Sex: "男"})
	// 删除一条数据
	//response, err := client.DeleteOne("student", "653f1f797ed5bb441885c00d")
	// 更新一条数据
	//response, err := client.UpdateOne("student", "653f1f797ed5bb441885c00d", map[string]interface{}{"$set": &Student{Name: "王华华", Age: 22, Sex: "女"}})
	// 查询一条数据
	//response, err := client.ReadOne("student", "653f1f797ed5bb441885c00d")
	// 查询全部数据（分页查询）
	//response, err := client.ReadAll("student", 1, 10)
	defer response.Body.Close()
	if err != nil {
		log.Printf("网络错误: %v", err)
		return
	}
	apiResponse := &Response{}
	if err := json.NewDecoder(response.Body).Decode(apiResponse); err != nil {
		log.Printf("解析JSON数据失败: %v", err)
		return
	}
	if apiResponse.Code != 200 {
		log.Printf("Code: %d", apiResponse.Code)
		log.Printf("Message: %s", apiResponse.Message)
		log.Printf("创建失败")
		return
	}
	log.Printf("Code: %d", apiResponse.Code)
	log.Printf("Message: %s", apiResponse.Message)
	dataBytes, _ := json.Marshal(apiResponse.Data)
	log.Printf("Data: %s", string(dataBytes))
	log.Printf("创建成功: %v", err)
}
