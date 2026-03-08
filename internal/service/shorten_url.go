package service

import (
	"bytes"
	"context"
	"crypto/rand"
	"math/big"

	"github.com/senn404/bookmark-managent/internal/repository"
)

const (
	urlLength = 9
)

type shortenURL struct {
	urlStorage repository.URLStorage
}

type ShortenURLService interface {
}

func NewShortenURLService(urlStorage repository.URLStorage) ShortenURLService {
	return &shortenURL{
		urlStorage: urlStorage,
	}
}

func (s *shortenURL) ShortenURL(ctx context.Context, url string) (string, error) {
	var urlShorten bytes.Buffer

	for i := 1; i <= urlLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		urlShorten.WriteByte(charset[randomIndex.Int64()])
	}

	s.urlStorage.StoreURL(ctx, urlShorten.String(), url)

	return urlShorten.String(), nil
}
