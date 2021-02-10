package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/psghahremani/url-shortener/adapters/driving/http/models"
	"github.com/psghahremani/url-shortener/domain/services"

	"github.com/labstack/echo/v4"
)

func ShortenUrl(urlShortenerService services.UrlShortenerService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Read request body into a model.
		shortenUrlRequest := models.ShortenUrlRequest{}
		err := ctx.Bind(&shortenUrlRequest)
		if err != nil {
			var response struct {
				Error string `json:"error"`
			}
			response.Error = err.Error()
			return ctx.JSON(http.StatusBadRequest, response)
		}

		// Check if URL is provided.
		if shortenUrlRequest.Url == nil {
			var response struct {
				Error string `json:"error"`
			}
			response.Error = "url cannot be empty"
			return ctx.JSON(http.StatusBadRequest, response)
		}

		// If expires_at is provided, convert it from Unix format to Go time format.
		var expiresAt *time.Time
		if shortenUrlRequest.ExpiresAt != nil {
			convertedTime := time.Unix(*shortenUrlRequest.ExpiresAt, 0)
			expiresAt = &convertedTime
		}

		// Shorten URL using domain methods.
		shortenedUrl, err := urlShortenerService.ShortenUrl(
			ctx.Request().Context(),
			*shortenUrlRequest.Url,
			expiresAt,
		)
		if err != nil {
			var response struct {
				Error string `json:"error"`
			}
			response.Error = fmt.Sprintf("cannot store shortened url: %s", err.Error())
			return ctx.JSON(http.StatusInternalServerError, response)
		}

		// Return shortened URL in HTTP response.
		return ctx.JSON(http.StatusOK, shortenedUrl)
	}
}

func RedirectUrl(urlShortenerService services.UrlShortenerService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Lookup for URL using handle.
		handle := ctx.Param("handle")
		shortenedUrl, err := urlShortenerService.LookupUrl(ctx.Request().Context(), handle)
		if err != nil {
			var response struct {
				Error string `json:"error"`
			}
			response.Error = fmt.Sprintf("cannot redirect url: %s", err.Error())
			return ctx.JSON(http.StatusInternalServerError, response)
		}

		// Redirect user to original URL.
		return ctx.Redirect(http.StatusTemporaryRedirect, shortenedUrl.OriginalUrl)
	}
}
