package main

import (
	"crypto/rand"
	"encoding/base64"
)

type URLMapping struct {
	ID          int
	OriginalURL string
	ShortCode   string
}

func generateShortCode() (string, error) {
	b := make([]byte, 6)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b)[:6], nil
}
