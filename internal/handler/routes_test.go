package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zuu-development/fullstack-examination-2024/internal/db"
)

func TestRegister(t *testing.T) {
	// Setup
	e := echo.New()
	dbInstance, err := db.NewMemory()
	require.NoError(t, err)
	err = db.Migrate(dbInstance)
	require.NoError(t, err)
	Register(e, dbInstance)

	// Test cases
	tests := []struct {
		name         string
		method       string
		target       string
		expectedCode int
	}{
		{"Health_Check", http.MethodGet, "/api/v1/healthz", http.StatusOK},
		{"Create_Todo_without_body", http.MethodPost, "/api/v1/todos", http.StatusBadRequest}, // Assuming no body is sent, should return BadRequest
		{"Get_all_Todos", http.MethodGet, "/api/v1/todos", http.StatusOK},
		{"Get_non-existent_Todo", http.MethodGet, "/api/v1/todos/1", http.StatusNotFound},       // Assuming no todo with id 1 exists
		{"Update_Todo_without_body", http.MethodPut, "/api/v1/todos/1", http.StatusNotFound},    // Assuming no body is sent, should return BadRequest
		{"Delete_non-existent_Todo", http.MethodDelete, "/api/v1/todos/1", http.StatusNotFound}, // Assuming no todo with id 1 exists
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.target, nil)
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			assert.Equal(t, tt.expectedCode, rec.Code)
		})
	}
}
