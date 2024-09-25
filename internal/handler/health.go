package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthHandler interface {
	Healthz(c echo.Context) error
}

type healthHandler struct {
}

func NewHealth() HealthHandler {
	return &healthHandler{}
}

func (t *healthHandler) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, time.Now())
}
