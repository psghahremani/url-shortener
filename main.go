package main

import (
	"fmt"
	"log"
	"net"

	"github.com/psghahremani/url-shortener/adapters/driven/mongodb"
	grpcAdapter "github.com/psghahremani/url-shortener/adapters/driving/grpc"
	"github.com/psghahremani/url-shortener/adapters/driving/grpc/models"
	"github.com/psghahremani/url-shortener/config"
	"github.com/psghahremani/url-shortener/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create a URL repository using configurations.
	urlRepository, err := mongodb.NewMongoRepository(
		config.Config.MongoDbConfig.ConnectionUrl,
		config.Config.MongoDbConfig.DatabaseName,
		config.Config.MongoDbConfig.Timeout,
	)
	if err != nil {
		log.Fatalf("cannot initialie URL repository: %s", err.Error())
	}

	// Create a URL shortener service, pass repository as a dependency.
	urlShortenerService := domain.NewUrlShortenerService(urlRepository)

	fmt.Println("yea")

	// Start gRPC server.
	urlShortenerGrpcServer := grpcAdapter.NewUrlShortenerGrpcServer(urlShortenerService)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Config.GrpcServer.Port))
	if err != nil {
		log.Fatalf("cannot initialie a URL shortener gRPC server: %s", err.Error())
	}

	server := grpc.NewServer()
	models.RegisterUrlShortenerServiceServer(server, urlShortenerGrpcServer)
	reflection.Register(server)

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %s", err.Error())
	}
}
