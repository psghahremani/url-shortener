package config

import (
	"fmt"
	"os"
	"strconv"
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
		Host:         getEnvOrDefault("MONGO_HOST", "172.17.0.2"),
		Port:         getEnvOrDefault("MONGO_PORT", "27017"),
		DatabaseName: getEnvOrDefault("MONGO_DATABASE_NAME", "url_shortener"),
	}

	mongoTimeout := getEnvOrDefault("MONGO_TIMEOUT", "30")
	mongoTimeoutInteger, err := strconv.Atoi(mongoTimeout)
	if err != nil {
		mongoTimeoutInteger = 30
	}
	Config.MongoDbConfig.Timeout = time.Duration(mongoTimeoutInteger) * time.Second
	Config.MongoDbConfig.ConnectionUrl = fmt.Sprintf(
		"mongodb://%s:%s",
		Config.MongoDbConfig.Host,
		Config.MongoDbConfig.Port,
	)
}
