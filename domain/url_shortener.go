package domain

import (
	"context"
	"time"

	"github.com/psghahremani/url-shortener/domain/models"
	"github.com/psghahremani/url-shortener/domain/ports"
	"github.com/psghahremani/url-shortener/domain/services"

	"github.com/matoous/go-nanoid/v2"
)

type urlShortener struct {
	urlRepository ports.UrlRepository
}

func NewUrlShortenerService(urlRepository ports.UrlRepository) services.UrlShortenerService {
	return &urlShortener{
		urlRepository: urlRepository,
	}
}

/* The following methods are the implementation of UrlShortenerService specified in domain services. */

func (urlShortener *urlShortener) ShortenUrl(ctx context.Context, originalUrl string, expiresAt *time.Time) (*models.ShortenedUrl, error) {
	// Create a handle for this URL.
	handle, err := gonanoid.New()
	if err != nil {
		return nil, err
	}

	// Make a shortened URL from the given URL.
	shortenedUrl := models.ShortenedUrl{
		Handle:      handle,
		OriginalUrl: originalUrl,
		CreatedAt:   time.Now(),
	}

	// If expiry date is not provided or it is invalid, expire URL in 24 hours.
	if expiresAt == nil || expiresAt.Sub(time.Now()) < 0 {
		shortenedUrl.ExpiresAt = time.Now().Add(24 * time.Hour)
	} else {
		shortenedUrl.ExpiresAt = *expiresAt
	}

	// Store the newly shortened URL in the repository.
	err = urlShortener.urlRepository.StoreUrl(ctx, shortenedUrl)
	if err != nil {
		return nil, err
	}
	return &shortenedUrl, nil
}

func (urlShortener *urlShortener) LookupUrl(ctx context.Context, handle string) (*models.ShortenedUrl, error) {
	// Lookup for URL in repository.
	return urlShortener.urlRepository.GetUrlByHandle(ctx, handle)
}
