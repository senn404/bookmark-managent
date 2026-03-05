package service

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"github.com/senn404/bookmark-managent/internal/config"
)

// TestMessRespon verifies that the HealthCheck service correctly builds its
// HealthStatus from configuration values. It tests two cases:
// 1. InstanceID is provided in config and used as-is.
// 2. InstanceID is empty, so a valid UUID is auto-generated.
func TestMessRespon(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name string

		setupConfg func() *config.Config

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

			expectedMess:       "OK",
			expectedServerName: "bookmark_service",
			expectedInstanceID: "",
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cfg := tc.setupConfg
			testSvc := NewHealthCheck(cfg())

			status := testSvc.GetStatus()

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
