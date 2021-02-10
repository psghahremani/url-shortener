package services

import (
	"context"
	"time"

	"github.com/psghahremani/url-shortener/domain/models"
)

type UrlShortenerCommands interface {
	ShortenUrl(context.Context, string, *time.Time) (*models.ShortenedUrl, error)
}

type UrlShortenerQueries interface {
	LookupUrl(context.Context, string) (*models.ShortenedUrl, error)
}

type UrlShortenerService interface {
	UrlShortenerCommands
	UrlShortenerQueries
}
