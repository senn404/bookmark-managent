package repository

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/redis/go-redis/v9"
	redisPkg "github.com/senn404/bookmark-managent/internal/pkg/redis"
)

func TestURLStorage(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name string

		setupMock func() *redis.Client
		inputCode string
		inputURL  string

		expectedErr error
		verifyFunc  func(ctx context.Context, r *redis.Client, inputCode, inputURL string)
	}{
		{
			name: "normal case",

			setupMock: func() *redis.Client {
				mock := redisPkg.InitMockRedis(t)
				return mock
			},
			inputURL:  "huanops.com",
			inputCode: "12345",

			expectedErr: nil,
			verifyFunc: func(ctx context.Context, r *redis.Client, inputCode, inputURL string) {
				res, err := r.Get(ctx, inputCode).Result()
				assert.Equal(t, nil, err)
				assert.Equal(t, inputURL, res)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := t.Context()
			redisClient := *tc.setupMock()

			urlStorage := NewURLStorage(&redisClient)
			err := urlStorage.StoreURL(ctx, tc.inputCode, tc.inputURL)
			if err == nil {
				tc.verifyFunc(ctx, &redisClient, tc.inputCode, tc.inputURL)
			}
		})
	}
}
