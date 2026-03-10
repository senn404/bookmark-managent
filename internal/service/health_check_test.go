package service

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"github.com/senn404/bookmark-managent/internal/config"
	"github.com/senn404/bookmark-managent/internal/repository/mocks"
	"github.com/stretchr/testify/mock"
)

// TestMessRespon verifies that the HealthCheck service correctly builds its
// HealthStatus from configuration values. It tests two cases:
// 1. InstanceID is provided in config and used as-is.
// 2. InstanceID is empty, so a valid UUID is auto-generated.
func TestMessRespon(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name string

		setupConfg            func() *config.Config
		setupRedisHealthCheck func() *mocks.HealthCheckRedis

		expectedMess       string
		expectedServerName string
		expectedInstanceID string
	}{
		{
			name: "case 1: read config?",

			setupConfg: func() *config.Config {
				return &config.Config{
					ServiceName: "bookmark_service",
					InstanceID:  "123456789",
				}
			},
			setupRedisHealthCheck: func() *mocks.HealthCheckRedis {
				mockHealtCheckRedis := mocks.NewHealthCheckRedis(t)
				mockHealtCheckRedis.On("HealthCheck", mock.Anything).Return(nil)
				return mockHealtCheckRedis
			},

			expectedMess:       "OK",
			expectedServerName: "bookmark_service",
			expectedInstanceID: "123456789",
		},
		{
			name: "case 2: create uuid?",

			setupConfg: func() *config.Config {
				return &config.Config{
					ServiceName: "bookmark_service",
					InstanceID:  "",
				}
			},
			setupRedisHealthCheck: func() *mocks.HealthCheckRedis {
				mockHealtCheckRedis := mocks.NewHealthCheckRedis(t)
				mockHealtCheckRedis.On("HealthCheck", mock.Anything).Return(nil)
				return mockHealtCheckRedis
			},

			expectedMess:       "OK",
			expectedServerName: "bookmark_service",
			expectedInstanceID: "",
		},
		{
			name: "case 3: test redis",

			setupConfg: func() *config.Config {
				return &config.Config{
					ServiceName: "bookmark_service",
					InstanceID:  "123456789",
				}
			},
			setupRedisHealthCheck: func() *mocks.HealthCheckRedis {
				mockHealtCheckRedis := mocks.NewHealthCheckRedis(t)
				mockHealtCheckRedis.On("HealthCheck", mock.Anything).Return(nil)
				return mockHealtCheckRedis
			},

			expectedMess:       "OK",
			expectedServerName: "bookmark_service",
			expectedInstanceID: "123456789",
		},
		{
			name: "case 4: redis fail",

			setupConfg: func() *config.Config {
				return &config.Config{
					ServiceName: "bookmark_service",
					InstanceID:  "123456789",
				}
			},
			setupRedisHealthCheck: func() *mocks.HealthCheckRedis {
				mockHealtCheckRedis := mocks.NewHealthCheckRedis(t)
				mockHealtCheckRedis.On("HealthCheck", mock.Anything).Return(errors.New("fail"))
				return mockHealtCheckRedis
			},

			expectedMess:       "Internal Server Error",
			expectedServerName: "bookmark_service",
			expectedInstanceID: "123456789",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cfg := tc.setupConfg

			mockHealtCheckRedis := tc.setupRedisHealthCheck()

			testSvc := NewHealthCheck(cfg(), mockHealtCheckRedis)

			status := testSvc.GetStatus(t.Context())

			assert.Equal(t, tc.expectedMess, status.Message)
			assert.Equal(t, tc.expectedServerName, status.ServiceName)

			if tc.expectedInstanceID == "" {
				_, err := uuid.Parse(status.InstanceId)
				assert.Equal(t, nil, err)
				assert.NotEqual(t, "", status.InstanceId)
			} else {
				assert.Equal(t, tc.expectedInstanceID, status.InstanceId)
			}
		})
	}
}
