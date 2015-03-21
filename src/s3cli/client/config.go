package client

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
)

type Config struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	BucketName      string `json:"bucket_name"`

	Host   string `json:"host"`
	Port   int    `json:"port"` // 0 means no custom port
	UseSSL bool   `json:"use_ssl"`
}

func NewConfigFromPath(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	config := Config{UseSSL: true, Port: 443}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c Config) Scheme() string {
	if c.UseSSL {
		return "https"
	}

	return "http"
}

func (c Config) HostWithPort() string {
	host := "s3.amazonaws.com"
	if c.Host != "" {
		host = c.Host
	}

	portSuffix := ""
	if c.Port != 443 {
		portSuffix = ":" + strconv.Itoa(c.Port)
	}

	return host + portSuffix
}
