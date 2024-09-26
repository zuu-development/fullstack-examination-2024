package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/zuu-development/fullstack-examination-2024/internal/repository"
	"github.com/zuu-development/fullstack-examination-2024/internal/service"
	"gorm.io/gorm"
)

// Register registers the routes for the application.
func Register(e *echo.Echo, db *gorm.DB) {
	healthHandler := NewHealth()
	e.GET("/healthz", healthHandler.Healthz)

	api := e.Group("/api")

	repository := repository.NewTodo(db)
	service := service.NewTodo(repository)
	todoHandler := NewTodo(service)
	todo := api.Group("/todo")
	{
		todo.POST("", todoHandler.Create)
		todo.GET("", todoHandler.FindAll)
		todo.GET("/:id", todoHandler.Find)
		todo.PUT("/:id", todoHandler.Update)
		todo.DELETE("/:id", todoHandler.Delete)
	}
}
