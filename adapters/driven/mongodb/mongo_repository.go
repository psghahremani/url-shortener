package mongodb

import (
	"context"
	"time"

	"github.com/psghahremani/url-shortener/adapters/driven/mongodb/models"
	domainModels "github.com/psghahremani/url-shortener/domain/models"
	"github.com/psghahremani/url-shortener/domain/ports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client       *mongo.Client
	databaseName string
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

func NewMongoRepository(mongoConnectionUrl, databaseName string, timeout time.Duration) (ports.UrlRepository, error) {
	// Create a MongoDB client using the given URL.
	client, err := newMongoClient(mongoConnectionUrl, timeout)
	if err != nil {
		return nil, err
	}

	// Create and return repository.
	repository := &mongoRepository{
		client:       client,
		databaseName: databaseName,
	}
	return repository, nil
}

/* The following methods are the implementation of UrlRepository specified in domain ports. */

// TODO: Automatically remove expired documents.
func (mongoRepository *mongoRepository) StoreUrl(ctx context.Context, shortenedUrl domainModels.ShortenedUrl) error {
	// Convert domain model to repository model.
	shortenedUrlDto := models.ShortenedUrl{
		Handle:      shortenedUrl.Handle,
		OriginalUrl: shortenedUrl.OriginalUrl,
		CreatedAt:   shortenedUrl.CreatedAt,
		ExpiresAt:   shortenedUrl.ExpiresAt,
	}

	// Insert the new shortened URL into the database.
	_, err := mongoRepository.client.Database(mongoRepository.databaseName).Collection("urls").InsertOne(ctx, shortenedUrlDto)

	return err
}

func (mongoRepository *mongoRepository) GetUrlByHandle(ctx context.Context, handle string) (*domainModels.ShortenedUrl, error) {
	// Prepare a query document.
	query := bson.D{
		{"handle", handle},
	}

	// Execute query and fetch the result.
	var result models.ShortenedUrl
	err := mongoRepository.client.Database(mongoRepository.databaseName).Collection("urls").FindOne(ctx, query).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Convert result to domain model.
	shortenedUrlDto := domainModels.ShortenedUrl{
		Handle:      result.Handle,
		OriginalUrl: result.OriginalUrl,
		CreatedAt:   result.CreatedAt,
		ExpiresAt:   result.ExpiresAt,
	}
	return &shortenedUrlDto, nil
}
