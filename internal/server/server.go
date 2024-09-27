// Package server provides the API server for the application.
package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/zuu-development/fullstack-examination-2024/internal/common"
	"github.com/zuu-development/fullstack-examination-2024/internal/db"
	"github.com/zuu-development/fullstack-examination-2024/internal/handler"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"gorm.io/gorm"
)

// TodoServer is the API server for Todo
type TodoServer struct {
	port   int
	engine *echo.Echo
	log    *log.Entry
	db     *gorm.DB
}

// TodoServerOpts is the options for the TodoServer
type TodoServerOpts struct {
	ListenPort int
	Config     model.Config
}

// NewServer returns a new instance of the Todo API server
func NewServer(opts TodoServerOpts) *TodoServer {
	logger := log.NewEntry(log.StandardLogger())
	log.SetFormatter(&log.JSONFormatter{})

	dbInstance, err := db.New(opts.Config.SQLite.DBFilename)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	engine := echo.New()
	engine.HideBanner = true
	engine.HidePort = true

	handler.Register(engine, dbInstance)
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	engine.Use(requestLogger())

	s := &TodoServer{
		port:   opts.ListenPort,
		engine: engine,
		log:    logger,
		db:     dbInstance,
	}
	return s
}

// Run starts the Todo API server
func (s *TodoServer) Run() error {
	log.Infof("todo %s serving on port %d", common.GetVersion(), s.port)
	return s.engine.Start(fmt.Sprintf(":%d", s.port))
}

// Shutdown stops the Todo API server
func (s *TodoServer) Shutdown(ctx context.Context) error {
	log.Infof("shuting down todo %s serving on port %d", common.GetVersion(), s.port)
	return s.engine.Shutdown(ctx)
}
