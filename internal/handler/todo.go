package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"github.com/zuu-development/fullstack-examination-2024/internal/service"
)

// TodoHandler is the request handler for the todo endpoint.
type TodoHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	Find(c echo.Context) error
	FindAll(c echo.Context) error
}

type todoHandler struct {
	Handler
	service service.Todo
}

// NewTodo returns a new instance of the todo handler.
func NewTodo(s service.Todo) TodoHandler {
	return &todoHandler{service: s}
}

// CreateRequest is the request parameter for creating a new todo
type CreateRequest struct {
	Task string `json:"task" validate:"required"`
}

func (t *todoHandler) Create(c echo.Context) error {
	var req CreateRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	todo, err := t.service.Create(req.Task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, Response{Data: todo})
}

// UpdateRequest is the request parameter for updating a todo
type UpdateRequest struct {
	ID     int          `param:"id" validate:"required"`
	Task   string       `json:"task,omitempty"`
	Status model.Status `json:"status,omitempty"`
}

func (t *todoHandler) Update(c echo.Context) error {
	var req UpdateRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err})
	}

	todo, err := t.service.Update(req.ID, req.Task, req.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: err})
	}

	return c.JSON(http.StatusNoContent, Response{Data: todo})
}

// DeleteRequest is the request parameter for deleting a todo
type DeleteRequest struct {
	ID int `param:"id" validate:"required"`
}

func (t *todoHandler) Delete(c echo.Context) error {
	var req DeleteRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err})
	}
	if err := t.service.Delete(req.ID); err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: err})
	}
	return c.JSON(http.StatusNoContent, Response{})
}

// FindRequest is the request parameter for finding a todo
type FindRequest struct {
	ID int `param:"id" validate:"required"`
}

func (t *todoHandler) Find(c echo.Context) error {
	var req FindRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, Response{Error: err})
	}
	res, err := t.service.Find(req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: err})
	}
	return c.JSON(http.StatusOK, Response{Data: res})
}

func (t *todoHandler) FindAll(c echo.Context) error {
	res, err := t.service.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{Error: err})
	}
	return c.JSON(http.StatusOK, Response{Data: res})
}
