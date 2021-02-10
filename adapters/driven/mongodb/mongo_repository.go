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

func newMongoClient(ctx context.Context, mongoConnectionUrl string) (*mongo.Client, error) {
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
	// Make a context based on the given timeout.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create a MongoDB client using the given URL.
	client, err := newMongoClient(ctx, mongoConnectionUrl)
	if err != nil {
		return nil, err
	}

	// Create an index on "expires_at" key, and automatically expire documents with this key.
	var expireAfterSeconds int32 = 0
	_, err = client.Database(databaseName).Collection("urls").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{
			"expires_at": 1,
		},
		Options: &options.IndexOptions{
			ExpireAfterSeconds: &expireAfterSeconds,
		},
	})
	if err != nil {
		return nil, err
	}

	// Create a unique index on "original_url".
	unique := true
	_, err = client.Database(databaseName).Collection("urls").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{
			"original_url": 1,
		},
		Options: &options.IndexOptions{
			Unique: &unique,
		},
	})
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
