package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/senn404/bookmark-managent/internal/service"
	"github.com/senn404/bookmark-managent/internal/service/mocks"
)

// TestHealthCheckHandler_HealthCheck verifies that the HealthCheck handler
// correctly returns the service health status as a JSON response.
// It uses a mock HealthCheck service to isolate the handler logic.
func TestHealthCheckHandler_HealthCheck(t *testing.T) {
	t.Parallel()

	testCase := []struct {
		name string

		setupRequest func(ctx *gin.Context)
		setupMockSvc func() *mocks.HealthCheck

		expectedStatus int
		expectedResp   string
	}{
		{
			name: "success",

			setupRequest: func(ctx *gin.Context) {
				ctx.Request = httptest.NewRequest(http.MethodGet, "/health-check", nil)
			},
			setupMockSvc: func() *mocks.HealthCheck {
				svcMock := mocks.NewHealthCheck(t)
				svcMock.On("GetStatus").Return(service.HealthStatus{
					Message:     "OK",
					ServiceName: "bookmark_service",
					InstanceId:  "test-123-123-123",
				})
				return svcMock
			},

			expectedStatus: http.StatusOK,
			expectedResp:   `{"message":"OK","service_name":"bookmark_service","instance_id":"test-123-123-123"}`,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			gc, _ := gin.CreateTestContext(rec)

			tc.setupRequest(gc)
			mockSvc := tc.setupMockSvc()

			testHandler := NewHealthCheckHandler(mockSvc)
			testHandler.HealthCheck(gc)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResp, rec.Body.String())
		})
	}
}
