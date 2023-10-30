package t1y

import (
	"net/http"
)

type Config struct {
	Domain    string
	AppID     int
	APIKey    string
	SecretKey string
}

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

func (config *Config) Read(table, id string) (*http.Response, error) {
	return config.sendRequest("GET", table+"/"+id, nil)
}
