package t1y

import (
	"fmt"
	"net/http"
)

func Init(domain string, appID int, apiKey, secretKey string) *Config {
	return &Config{
		Domain:    domain,
		AppID:     appID,
		APIKey:    apiKey,
		SecretKey: secretKey,
	}
}

func (config *Config) CreateOne(table string, data interface{}) (*http.Response, error) {
	return config.sendRequest("POST", table, data)
}

func (config *Config) DeleteOne(table string, id string) (*http.Response, error) {
	return config.sendRequest("DELETE", table+"/"+id, nil)
}

func (config *Config) UpdateOne(table, id string, data interface{}) (*http.Response, error) {
	return config.sendRequest("PUT", table+"/"+id, data)
}

func (config *Config) ReadOne(table, id string) (*http.Response, error) {
	return config.sendRequest("GET", table+"/"+id, nil)
}

func (config *Config) ReadAll(table string, page, size int) (*http.Response, error) {
	return config.sendRequest("GET", fmt.Sprintf("%s?page=%d&size=%d", table, page, size), nil)
}
