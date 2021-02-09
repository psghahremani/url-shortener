package ports

import (
	"context"

	"github.com/psghahremani/url-shortener/domain/models"
)

type UrlRepositoryCommands interface {
	StoreUrl(context.Context, models.ShortenedUrl) error
}

type UrlRepositoryQueries interface {
	GetUrlByHandle(context.Context, string) (*models.ShortenedUrl, error)
}

type UrlRepository interface {
	UrlRepositoryCommands
	UrlRepositoryQueries
}
