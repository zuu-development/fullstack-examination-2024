package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/go-cmp/cmp"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zuu-development/fullstack-examination-2024/internal/db"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/repository"
	"github.com/zuu-development/fullstack-examination-2024/internal/service"
)

func TestTodoHandler_Create(t *testing.T) {
	type want struct {
		StatusCode int
		Response   []byte
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	dbInstance, err := db.NewMemory()
	require.NoError(t, err)
	err = db.Migrate(dbInstance)
	require.NoError(t, err)
	repository := repository.NewTodo(dbInstance)
	service := service.NewTodo(repository)
	handler := NewTodo(service)

	tests := []struct {
		name       string
		createBody string
		want       want
		wantErr    bool
	}{
		{
			name:       "successful_create",
			createBody: `{"task":"Created Task"}`,
			want: want{
				StatusCode: http.StatusCreated,
				Response:   []byte(`{"data":{"Task":"Created Task","Status":"created"}}`),
			},
		},
		{
			name:       "successful_create_but_with_ignore_status",
			createBody: `{"task":"Created Task", "status":"done"}`,
			want: want{
				StatusCode: http.StatusCreated,
				Response:   []byte(`{"data":{"Task":"Created Task","Status":"created"}}`),
			},
		},
		{
			name:       "invalid_request_body",
			createBody: `{"task":1}`,
			want: want{
				StatusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			req := httptest.NewRequest(http.MethodPost, "/dummy/target", bytes.NewReader([]byte(tt.createBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos")

			// Execute
			require.NoError(t, handler.Create(c))

			// Assert
			assert.Equal(t, tt.want.StatusCode, rec.Code)

			if tt.want.Response == nil {
				return
			}
			got := rec.Body.Bytes()

			opts := []cmp.Option{
				cmpTransformJSON(t),
				ignoreMapEntires(map[string]any{"CreatedAt": 1, "UpdatedAt": 1, "ID": 1}),
			}
			if diff := cmp.Diff(got, tt.want.Response, opts...); diff != "" {
				t.Errorf("return value mismatch (-got +want):\n%s", diff)
				t.Logf("got:\n%s", string(got))
			}
		})
	}
}

func TestTodoHandler_Update(t *testing.T) {
	type want struct {
		StatusCode int
		Response   []byte
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	dbInstance, err := db.NewMemory()
	require.NoError(t, err)
	err = db.Migrate(dbInstance)
	require.NoError(t, err)
	repository := repository.NewTodo(dbInstance)
	service := service.NewTodo(repository)
	handler := NewTodo(service)

	tests := []struct {
		name       string
		createBody string
		updateBody string
		updateID   string
		want       want
		wantErr    bool
	}{
		{
			name:       "successful_update",
			createBody: `{"task":"Updated Task"}`,
			updateBody: `{"task":"Updated Task","status":"done"}`,
			want: want{
				StatusCode: http.StatusOK,
				Response:   []byte(`{"data":{"Task":"Updated Task","Status":"done"}}`),
			},
		},
		{
			name:       "not_found_record",
			updateID:   "-1",
			updateBody: `{"task":"Updated Task","status":"done"}`,
			want: want{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:       "invalid_request_body",
			updateBody: `{"task":1}`,
			want: want{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			name:       "invalid_request_parameter",
			updateID:   "invalid",
			updateBody: `{"task":"Updated Task","status":"done"}`,
			want: want{
				StatusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			id := "1"
			if tt.updateID != "" {
				id = tt.updateID
			} else if tt.createBody != "" {
				id = strconv.Itoa(createTask(t, e, handler, tt.createBody))
			}

			req := httptest.NewRequest(http.MethodPut, "/dummy/target", bytes.NewReader([]byte(tt.updateBody)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos/:id")
			c.SetParamNames("id")
			c.SetParamValues(id)

			// Execute
			require.NoError(t, handler.Update(c))

			// Assert
			assert.Equal(t, tt.want.StatusCode, rec.Code)

			if tt.want.Response == nil {
				return
			}
			got := rec.Body.Bytes()

			opts := []cmp.Option{
				cmpTransformJSON(t),
				ignoreMapEntires(map[string]any{"CreatedAt": 1, "UpdatedAt": 1, "ID": 1}),
			}
			if diff := cmp.Diff(got, tt.want.Response, opts...); diff != "" {
				t.Errorf("return value mismatch (-got +want):\n%s", diff)
				t.Logf("got:\n%s", string(got))
			}
		})
	}
}

func TestTodoHandler_Delete(t *testing.T) {
	type want struct {
		StatusCode int
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	dbInstance, err := db.NewMemory()
	require.NoError(t, err)
	err = db.Migrate(dbInstance)
	require.NoError(t, err)
	repository := repository.NewTodo(dbInstance)
	service := service.NewTodo(repository)
	handler := NewTodo(service)

	tests := []struct {
		name       string
		createBody string
		deleteID   string
		want       want
		wantErr    bool
	}{
		{
			name:       "successful_delete",
			createBody: `{"task":"Deleted Task"}`,
			want: want{
				StatusCode: http.StatusNoContent,
			},
		},
		{
			name:     "not_found_record",
			deleteID: "-1",
			want: want{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:     "invalid_request_parameter",
			deleteID: "invalid",
			want: want{
				StatusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			id := "1"
			if tt.deleteID != "" {
				id = tt.deleteID
			} else if tt.createBody != "" {
				id = strconv.Itoa(createTask(t, e, handler, tt.createBody))
			}

			req := httptest.NewRequest(http.MethodDelete, "/dummy/target", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos/:id")
			c.SetParamNames("id")
			c.SetParamValues(id)

			// Execute
			require.NoError(t, handler.Delete(c))

			// Assert
			assert.Equal(t, tt.want.StatusCode, rec.Code)
		})
	}
}

func TestTodoHandler_Find(t *testing.T) {
	type want struct {
		StatusCode int
		Response   []byte
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	dbInstance, err := db.NewMemory()
	require.NoError(t, err)
	err = db.Migrate(dbInstance)
	require.NoError(t, err)
	repository := repository.NewTodo(dbInstance)
	service := service.NewTodo(repository)
	handler := NewTodo(service)

	tests := []struct {
		name       string
		createBody string
		findID     string
		want       want
		wantErr    bool
	}{
		{
			name:       "successful_find",
			createBody: `{"task":"Found Task"}`,
			want: want{
				StatusCode: http.StatusOK,
				Response:   []byte(`{"data":{"Task":"Found Task","Status":"created"}}`),
			},
		},
		{
			name:   "not_found",
			findID: "-1",
			want: want{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:   "invalid_request_parameter",
			findID: "invalid",
			want: want{
				StatusCode: http.StatusBadRequest,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			id := "1"
			if tt.findID != "" {
				id = tt.findID
			} else if tt.createBody != "" {
				id = strconv.Itoa(createTask(t, e, handler, tt.createBody))
			}

			req := httptest.NewRequest(http.MethodGet, "/dummy/target", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/todos/:id")
			c.SetParamNames("id")
			c.SetParamValues(id)

			// Execute
			require.NoError(t, handler.Find(c))

			// Assert
			assert.Equal(t, tt.want.StatusCode, rec.Code)

			if tt.want.Response == nil {
				return
			}
			got := rec.Body.Bytes()

			opts := []cmp.Option{
				cmpTransformJSON(t),
				ignoreMapEntires(map[string]any{"CreatedAt": 1, "UpdatedAt": 1, "ID": 1}),
			}
			if diff := cmp.Diff(got, tt.want.Response, opts...); diff != "" {
				t.Errorf("return value mismatch (-got +want):\n%s", diff)
				t.Logf("got:\n%s", string(got))
			}
		})
	}
}

func TestTodoHandler_FindAll(t *testing.T) {
	t.Skip("Not implemented")
}

func createTask(t *testing.T, e *echo.Echo, handler TodoHandler, body string) int {
	req := httptest.NewRequest(http.MethodPost, "/todos", bytes.NewReader([]byte(body)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := handler.Create(c)
	require.NoError(t, err)

	type response struct {
		Data   model.Todo
		Status string
	}

	var res response
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	require.NoError(t, err, "Failed to json.Unmarshal err: %s", err)
	require.NotEmpty(t, res.Data.ID)

	return res.Data.ID
}
