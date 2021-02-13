package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/psghahremani/url-shortener/adapters/driving/grpc/models"
	"github.com/psghahremani/url-shortener/domain/services"
)

type UrlShortenerGrpcServer struct {
	models.UnimplementedUrlShortenerServiceServer
	urlShortenerService services.UrlShortenerService
}

func NewUrlShortenerGrpcServer(urlShortenerService services.UrlShortenerService) *UrlShortenerGrpcServer {
	return &UrlShortenerGrpcServer{
		urlShortenerService: urlShortenerService,
	}
}

func (urlShortenerGrpcServer UrlShortenerGrpcServer) ShortenUrl(ctx context.Context, request *models.ShortenUrlRequest) (*models.ShortenedUrl, error) {
	if request.GetUrl() == "" {
		return nil, errors.New("url must be provided")
	}

	var expiresAt *time.Time

	if request.GetExpiresAt() != 0 {
		temp := time.Unix(request.GetExpiresAt(), 0)
		expiresAt = &temp
	}

	shortenedUrl, err := urlShortenerGrpcServer.urlShortenerService.ShortenUrl(ctx, request.GetUrl(), expiresAt)
	if err != nil {
		return nil, err
	}

	response := models.ShortenedUrl{
		Handle:      shortenedUrl.Handle,
		OriginalUrl: shortenedUrl.OriginalUrl,
		ExpiresAt:   shortenedUrl.ExpiresAt.Unix(),
		CreatedAt:   shortenedUrl.CreatedAt.Unix(),
	}

	return &response, nil
}

func (urlShortenerGrpcServer UrlShortenerGrpcServer) LookupUrl(ctx context.Context, request *models.LookupUrlRequest) (*models.ShortenedUrl, error) {
	if request.GetHandle() == "" {
		return nil, errors.New("handle must be provided")
	}

	shortenedUrl, err := urlShortenerGrpcServer.urlShortenerService.LookupUrl(ctx, request.GetHandle())
	if err != nil {
		return nil, err
	}

	response := models.ShortenedUrl{
		Handle:      shortenedUrl.Handle,
		OriginalUrl: shortenedUrl.OriginalUrl,
		ExpiresAt:   shortenedUrl.ExpiresAt.Unix(),
		CreatedAt:   shortenedUrl.CreatedAt.Unix(),
	}

	return &response, nil
}
