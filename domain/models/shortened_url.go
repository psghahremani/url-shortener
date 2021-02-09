package models

import (
	"time"
)

type ShortenedUrl struct {
	Handle      string
	OriginalUrl string
	CreatedAt   time.Time
	ExpiresAt   time.Time
}
