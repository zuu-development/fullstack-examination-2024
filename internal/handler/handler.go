// Package handler provides request handling functionality for the application.
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler is the request handler for the application.
type Handler struct{}

// MustBind はリクエストのバインドとバリデーションを行います。
func (h Handler) MustBind(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
