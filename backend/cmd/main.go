package main

import (
	"crypto/rand"
	"database/sql"
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

const (
	dbConnStr = ""
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS url_mappings (
			id SERIAL PRIMARY KEY
			original_url TEXT NOT NUL
			short_code VARCHAR(10) NOT NULL UNIQUE
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
		`)

	return db, err
}
