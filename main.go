package main

import (
	"fmt"
	"log"

	"github.com/psghahremani/url-shortener/adapters/driven/mongodb"
	"github.com/psghahremani/url-shortener/adapters/driving/http"
	"github.com/psghahremani/url-shortener/config"
	"github.com/psghahremani/url-shortener/domain"
)

func main() {
	urlRepository, err := mongodb.NewMongoRepository(
		config.Config.MongoDbConfig.ConnectionUrl,
		config.Config.MongoDbConfig.DatabaseName,
		config.Config.MongoDbConfig.Timeout,
	)
	if err != nil {
		log.Fatalf("cannot initialie URL repository: %s", err.Error())
	}

	urlShortenerService := domain.NewUrlShortenerService(urlRepository)

	e := http.NewEchoServer(urlShortenerService)
	err = e.Start(fmt.Sprintf(":%s", config.Config.HttpServer.Port))
	if err != nil {
		log.Fatalf("cannot start HTTP server: %s", err.Error())
	}
}
