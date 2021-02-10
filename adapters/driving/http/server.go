package http

import (
	"log"

	"github.com/psghahremani/url-shortener/domain/services"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(handler echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := handler(ctx)

		// Log error message if any errors occurred.
		if err != nil {
			httpError, ok := err.(*echo.HTTPError)
			if ok {
				log.Printf("%d %s\n", httpError.Code, httpError.Message)
			} else {
				log.Printf("%s\n", err.Error())
			}
		} else {
			// Log status, method and URL.
			log.Printf("%d %s %s\n", ctx.Response().Status, ctx.Request().Method, ctx.Request().RequestURI)
		}

		return err
	}
}

func NewEchoServer(urlShortenerService services.UrlShortenerService) *echo.Echo {
	// Create Echo server.
	e := echo.New()

	// User the logger middleware.
	e.Use(LoggerMiddleware)

	// Register routes.
	bindRoutes(e, urlShortenerService)

	return e
}
