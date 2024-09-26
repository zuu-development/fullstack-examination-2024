// Package handler provides request handling functionality for the application.
package handler

import (
	"github.com/labstack/echo/v4"
)

// Handler is the request handler for the application.
type Handler struct{}

// MustBind はリクエストのバインドとバリデーションを行います。
func (h Handler) MustBind(c echo.Context, req interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := validateStruct(req); err != nil {
		return err
	}
	return nil
}
