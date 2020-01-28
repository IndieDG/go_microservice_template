package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServeHealth(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	ServeHealth(response, request)

	t.Run("response should be OK", func(t *testing.T) {
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("body should be health", func(t *testing.T) {
		assert.Equal(t, "health", response.Body.String())
	})
}
