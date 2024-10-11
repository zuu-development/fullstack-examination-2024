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
	Task     string `json:"task" validate:"required"`
	Priority int    `json:"priority,omitempty"`
}

// Create handles the creation of a new todo item.
//
//	@Summary		Create a new todo
//	@Description	Creates a new todo item with the specified task and optional priority
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CreateRequest					true	"Todo creation request"
//	@Success		201		{object}	ResponseData{Data=model.Todo}	"Successfully created a new todo item"
//	@Failure		400		{object}	ResponseError					"Invalid request parameters"
//	@Failure		500		{object}	ResponseError					"Internal server error while creating the todo item"
//	@Router			/api/v1/todos [post]
func (t *todoHandler) Create(c echo.Context) error {
	var req CreateRequest
	if err := t.MustBind(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	todo, err := t.service.Create(req.Task, req.Priority)
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

// Update handles updating an existing todo item.
//
//	@Summary		Update a todo
//	@Description	Updates the details of a todo item with the specified ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"ID of the todo to be updated"
//	@Param			request	body		UpdateRequestBody				true	"Todo update request body"
//	@Success		200		{object}	ResponseData{Data=model.Todo}	"Successfully updated the todo item"
//	@Failure		400		{object}	ResponseError					"Invalid request parameters"
//	@Failure		404		{object}	ResponseError					"Todo item not found"
//	@Failure		500		{object}	ResponseError					"Internal server error while updating the todo item"
//	@Router			/api/v1/todos/{id} [put]
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

// Delete handles deleting an existing todo item.
//
//	@Summary		Delete a todo
//	@Description	Deletes a todo item with the specified ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"ID of the todo to be deleted"
//	@Success		204	"Successfully deleted the todo item"
//	@Failure		400	{object}	ResponseError	"Invalid request parameters"
//	@Failure		404	{object}	ResponseError	"Todo item not found"
//	@Failure		500	{object}	ResponseError	"Internal server error while deleting the todo item"
//	@Router			/api/v1/todos/{id} [delete]
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

// Find handles retrieving a todo item by its ID.
//
//	@Summary		Find a todo
//	@Description	Retrieves a todo item with the specified ID
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int								true	"ID of the todo to be retrieved"
//	@Success		200	{object}	ResponseData{Data=model.Todo}	"Successfully retrieved the todo item"
//	@Failure		400	{object}	ResponseError					"Invalid request parameters"
//	@Failure		404	{object}	ResponseError					"Todo item not found"
//	@Failure		500	{object}	ResponseError					"Internal server error while retrieving the todo item"
//	@Router			/api/v1/todos/{id} [get]
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

type findQueryParams struct {
	Task   string       `query:"task"`
	Status model.Status `query:"status" validate:"omitempty,oneof=created processing done"`
}

// FindAll Find todos by optional task and status
//
//	@Summary	Find todos by optional task and status
//	@Tags		todos
//	@Accept		json
//	@Produce	json
//	@Param		task	query		string							false	"Task to filter todos (supports partial matches)"
//	@Param		status	query		string							false	"Status to filter todos (must be one of: 'created', 'processing', or 'done')"
//	@Success	200		{array}		ResponseData{Data=model.Todo}	"Successfully retrieved todos"
//	@Failure	400		{object}	ResponseError					"Invalid query parameters"
//	@Failure	500		{object}	ResponseError					"Failed to fetch records from the database"
//	@Router		/api/v1/todos [get]
func (t *todoHandler) FindAll(c echo.Context) error {
	findQueryParams := findQueryParams{}

	err := t.MustBind(c, &findQueryParams)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			ResponseError{Errors: []Error{{Code: errors.CodeBadRequest, Message: err.Error()}}})
	}

	res, err := t.service.FindAll(findQueryParams.Task, findQueryParams.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			ResponseError{Errors: []Error{{Code: errors.CodeInternalServerError, Message: err.Error()}}})
	}
	return c.JSON(http.StatusOK, ResponseData{Data: res})
}
