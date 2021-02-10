package models

type ShortenUrlRequest struct {
	Url       *string `json:"url"`
	ExpiresAt *int64 `json:"expires_at"`
}
