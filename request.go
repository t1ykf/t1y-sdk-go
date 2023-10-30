package t1y

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (config *Config) sendRequest(method, table string, data interface{}) (*http.Response, error) {
	index := strings.Index(table, "?page=")
	_table := table
	if index != -1 {
		_table = table[:index]
	}
	url := fmt.Sprintf("%s/v5/classes/%s", config.Domain, table)

	var jsonBytes []byte
	if data != nil {
		jsonBytes, _ = json.Marshal(data)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	nonceStr := UUID()
	timestamp := time.Now().Unix()
	req.Header.Add("X-T1Y-Application-ID", fmt.Sprintf("%d", config.AppID))
	req.Header.Add("X-T1Y-Api-Key", config.APIKey)
	req.Header.Add("X-T1Y-Safe-NonceStr", nonceStr)
	req.Header.Add("X-T1Y-Safe-Timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Add("X-T1Y-Safe-Sign", MD5(fmt.Sprintf("%s%d%s%s%d%s", fmt.Sprintf("/v5/classes/%s", _table), config.AppID, config.APIKey, nonceStr, timestamp, config.SecretKey)))

	return client.Do(req)
}
