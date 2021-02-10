package config

import (
	"fmt"
	"os"
	"time"
)

var Config = struct {
	MongoDbConfig MongoDbConfig
}{}

type MongoDbConfig struct {
	Host          string
	Port          string
	DatabaseName  string
	Timeout       time.Duration
	ConnectionUrl string
}

func getEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func init() {
	Config.MongoDbConfig = MongoDbConfig{
		Host:         "172.17.0.1",
		Port:         "27017",
		Timeout:      30 * time.Second,
		DatabaseName: "url_shortener",
	}
	Config.MongoDbConfig.ConnectionUrl = fmt.Sprintf(
		"mongodb://%s:%s",
		Config.MongoDbConfig.Host,
		Config.MongoDbConfig.Port,
	)
}
