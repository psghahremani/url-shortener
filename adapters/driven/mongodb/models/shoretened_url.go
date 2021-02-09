package models

import (
	"time"
)

type ShortenedUrl struct {
	Handle      string    `bson:"handle"`
	OriginalUrl string    `bson:"original_url"`
	CreatedAt   time.Time `bson:"created_at"`
	ExpiresAt   time.Time `bson:"expires_at"`
}
