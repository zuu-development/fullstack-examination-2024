package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zuu-development/fullstack-examination-2024/internal/errors"
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

// @Summary	Create a new todo
// @Tags		todos
// @Accept		json
// @Produce	json
// @Param		request	body		CreateRequest	true	"json"
// @Success	201		{object}	ResponseError{data=model.Todo}
// @Failure	400		{object}	ResponseError
// @Failure	500		{object}	ResponseError
// @Router		/todos [post]
func (t *todoHandler) Create(c echo.Context) error {
	var req CreateRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	todo, err := t.service.Create(req.Task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}

	return c.JSON(http.StatusCreated, ResponseData{Data: todo})
}

// UpdateRequest is the request parameter for updating a todo
type UpdateRequest struct {
	UpdateRequestBody
	UpdateRequestPath
}

// UpdateRequestBody is the request body for updating a todo
type UpdateRequestBody struct {
	Task   string       `json:"task,omitempty"`
	Status model.Status `json:"status,omitempty"`
}

// UpdateRequestPath is the request parameter for updating a todo
type UpdateRequestPath struct {
	ID int `param:"id" validate:"required"`
}

// @Summary	Update a todo
// @Tags		todos
// @Accept		json
// @Produce	json
// @Param		body	body		UpdateRequestBody	true	"body"
// @Param		path	path		UpdateRequestPath	false	"path"
// @Success	201		{object}	ResponseData{Data=model.Todo}
// @Failure	400		{object}	ResponseError
// @Failure	500		{object}	ResponseError
// @Router		/todos/:id [put]
func (t *todoHandler) Update(c echo.Context) error {
	var req UpdateRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	todo, err := t.service.Update(req.ID, req.Task, req.Status)
	if err != nil {
		if err == model.ErrNotFound {
			return c.JSON(http.StatusNotFound,
				ResponseError{Errors: []Error{{Code: errors.CodeNotFound, Message: "todo not found"}}})
		}
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}

	return c.JSON(http.StatusOK, ResponseData{Data: todo})
}

// DeleteRequest is the request parameter for deleting a todo
type DeleteRequest struct {
	ID int `param:"id" validate:"required"`
}

// @Summary	Delete a todo
// @Tags		todos
// @Param		path	path	DeleteRequest	false	"path"
// @Success	204
// @Failure	400	{object}	ResponseError
// @Failure	404	{object}	ResponseError
// @Failure	500	{object}	ResponseError
// @Router		/todos/:id [delete]
func (t *todoHandler) Delete(c echo.Context) error {
	var req DeleteRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	if err := t.service.Delete(req.ID); err != nil {
		if err == model.ErrNotFound {
			return c.JSON(http.StatusNotFound,
				ResponseError{Errors: []Error{{Code: errors.CodeNotFound, Message: "todo not found"}}})
		}
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.NoContent(http.StatusNoContent)
}

// FindRequest is the request parameter for finding a todo
type FindRequest struct {
	ID int `param:"id" validate:"required"`
}

// @Summary	Find a todo
// @Tags		todos
// @Param		path	path		FindRequest	false	"path"
// @Success	200		{object}	ResponseData{Data=model.Todo}
// @Failure	400		{object}	ResponseError
// @Failure	404		{object}	ResponseError
// @Failure	500		{object}	ResponseError
// @Router		/todos/:id [get]
func (t *todoHandler) Find(c echo.Context) error {
	var req FindRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	res, err := t.service.Find(req.ID)
	if err != nil {
		if err == model.ErrNotFound {
			return c.JSON(http.StatusNotFound,
				ResponseError{Errors: []Error{{Code: errors.CodeNotFound, Message: "todo not found"}}})
		}
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.JSON(http.StatusOK, ResponseData{Data: res})
}

// @Summary	Find all todos
// @Tags		todos
// @Success	200	{object}	ResponseData{Data=[]model.Todo}
// @Failure	500	{object}	ResponseError
// @Router		/todos [get]
func (t *todoHandler) FindAll(c echo.Context) error {
	res, err := t.service.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.JSON(http.StatusOK, ResponseData{Data: res})
}
