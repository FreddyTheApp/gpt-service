package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSecretMiddleware(t *testing.T) {
	// Arrange
	secret := "super-secret"
	os.Setenv("API_SECRET", secret)

	// Create a new gin engine
	r := gin.Default()
	handler := GinHandler{}
	r.Use(handler.SecretMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	testCases := []struct {
		name        string
		headerValue string
		wantStatus  int
	}{
		{"valid secret", secret, http.StatusOK},
		{"invalid secret", "not-the-secret", http.StatusUnauthorized},
		{"empty secret", "", http.StatusUnauthorized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			req.Header.Set("X-Secret", tc.headerValue)
			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			// Assert
			if resp.Code != tc.wantStatus {
				t.Errorf("want status %d; got %d", tc.wantStatus, resp.Code)
			}
		})
	}
}
