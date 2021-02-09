package mongodb

import (
	"context"
	"time"

	domainModels "github.com/psghahremani/url-shortener/domain/models"
	"github.com/psghahremani/url-shortener/domain/ports"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	database *mongo.Database
}

func newMongoClient(mongoConnectionUrl string, timeout time.Duration) (*mongo.Client, error) {
	// Make a context based on the given timeout.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create a client using the given URL.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionUrl))
	if err != nil {
		return nil, err
	}

	// Ping database to ensure connection.
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoConnectionUrl, database string, timeout time.Duration) (ports.UrlRepository, error) {
	// Create a MongoDB client using the given URL.
	client, err := newMongoClient(mongoConnectionUrl, timeout)
	if err != nil {
		return nil, err
	}

	// Create and return repository.
	repository := &mongoRepository{
		database: client.Database(database),
	}
	return repository, nil
}

/* The following methods are the implementation of UrlRepository specified in domain ports. */

func (mongoRepository *mongoRepository) StoreUrl(context.Context, domainModels.ShortenedUrl) error {
	return nil
}

func (mongoRepository *mongoRepository) GetUrlByHandle(context.Context, string) (*domainModels.ShortenedUrl, error) {
	return nil, nil
}
