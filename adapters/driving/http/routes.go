package http

import (
	"net/http"

	"github.com/psghahremani/url-shortener/domain/services"

	"github.com/labstack/echo/v4"
)

/* A default health check endpoint used to check service availability. */
func HealthCheck(ctx echo.Context) error {
	var response struct {
		Message string `json:"message"`
	}
	response.Message = "Healthy!"
	return ctx.JSON(http.StatusOK, response)
}

func bindRoutes(e *echo.Echo, urlShortenerService services.UrlShortenerService) *echo.Echo {
	// Bind health check endpoint.
	e.GET("/health", HealthCheck)

	// Bind API grouping version 1.
	v1 := e.Group("/v1")
	v1.POST("/shorten", ShortenUrl(urlShortenerService))
	v1.GET("/:handle", RedirectUrl(urlShortenerService))

	return e
}
