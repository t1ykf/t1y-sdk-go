package t1y

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (sdk *Config) sendRequest(method, table string, data interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s/v5/classes/%s", sdk.Domain, table)

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
	req.Header.Add("X-T1Y-Application-ID", fmt.Sprintf("%d", sdk.AppID))
	req.Header.Add("X-T1Y-Api-Key", sdk.APIKey)
	req.Header.Add("X-T1Y-Safe-NonceStr", nonceStr)
	req.Header.Add("X-T1Y-Safe-Timestamp", fmt.Sprintf("%d", timestamp))
	req.Header.Add("X-T1Y-Safe-Sign", MD5(fmt.Sprintf("%s%d%s%s%d%s", url, sdk.AppID, sdk.APIKey, nonceStr, timestamp, sdk.SecretKey)))

	return client.Do(req)
}
