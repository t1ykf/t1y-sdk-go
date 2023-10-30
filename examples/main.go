package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/t1ykf/t1y-sdk-go"
)

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  int    `json:"sex"`
}

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	// 初始化 SDK 配置
	config := t1y.Init("http://dev.t1y.net/api", 1001, "", "")

	// 创建一条记录
	response, err := config.CreateOne("test", &Test{})
	if err != nil {
		fmt.Println("Error:", err.Error())
		//return
	}

	// 处理响应
	apiResponse := &APIResponse{}
	err = parseJSONResponse(response, apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		//return
	}

	fmt.Printf("Code: %d\n", apiResponse.Code)
	fmt.Printf("Message: %s\n", apiResponse.Message)

	// 打印 data 字段，你可以根据实际情况进一步处理 data 数据
	if apiResponse.Data != nil {
		dataBytes, _ := json.Marshal(apiResponse.Data)
		fmt.Printf("Data: %s\n", string(dataBytes))
	}
}

func parseJSONResponse(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return json.NewDecoder(response.Body).Decode(target)
	} else {
		return fmt.Errorf("HTTP request failed. Status Code: %d", response.StatusCode)
	}
}
