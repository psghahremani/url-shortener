package services

import (
	"context"

	"github.com/psghahremani/url-shortener/domain/models"
)

type UrlShortenerCommands interface {
	ShortenUrl(context.Context, string) (*models.ShortenedUrl, error)
}

type UrlShortenerQueries interface {
	LookupUrl(context.Context, string) (*models.ShortenedUrl, error)
}

type UrlShortenerService interface {
	UrlShortenerCommands
	UrlShortenerQueries
}
