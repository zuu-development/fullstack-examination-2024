package handler

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

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
