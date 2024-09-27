package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestRequest struct {
	Name string `json:"name" validate:"required"`
}

func TestMustBind(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	h := Handler{}

	tests := []struct {
		name           string
		requestBody    string
		expectedError  string
		expectedStatus int
	}{
		{
			name:        "Valid_request",
			requestBody: `{"name": "test"}`,
		},
		{
			name:           "Invalid_JSON",
			requestBody:    `{"name": "test"`,
			expectedError:  "unexpected EOF",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Validation_error",
			requestBody:    `{}`,
			expectedError:  "required",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.requestBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			var testReq TestRequest
			err := h.MustBind(c, &testReq)

			if tt.expectedError == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			if tt.expectedStatus == 0 {
				return
			}
			httpError, ok := err.(*echo.HTTPError)
			require.True(t, ok)
			assert.Equal(t, tt.expectedStatus, httpError.Code)
			assert.Contains(t, httpError.Error(), tt.expectedError)
		})
	}
}
