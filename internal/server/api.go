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

// todoAPIServer is the API server for Todo
type todoAPIServer struct {
	port   int
	engine *echo.Echo
	log    *log.Entry
	db     *gorm.DB
}

// TodoAPIServerOpts is the options for the TodoAPIServer
type TodoAPIServerOpts struct {
	ListenPort int
	Config     model.Config
}

// NewAPI returns a new instance of the Todo API server
func NewAPI(opts TodoAPIServerOpts) (Server, error) {
	logger := log.NewEntry(log.StandardLogger())
	log.SetFormatter(&log.JSONFormatter{})

	dbInstance, err := db.New(opts.Config.SQLite.DBFilename)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	engine := echo.New()
	engine.HideBanner = true
	engine.HidePort = true

	handler.Register(engine, dbInstance)

	allowOrigins := []string{opts.Config.UI.URL}
	if opts.Config.SwaggerServer.Enable {
		allowOrigins = append(allowOrigins, fmt.Sprintf("http://localhost:%d", opts.Config.SwaggerServer.Port))
	}
	log.Info("CORS allowed origins: ", allowOrigins)
	engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	engine.Use(requestLogger())

	s := &todoAPIServer{
		port:   opts.ListenPort,
		engine: engine,
		log:    logger,
		db:     dbInstance,
	}
	return s, nil
}

func (s *todoAPIServer) Name() string {
	return "todoAPIServer"
}

// Run starts the Todo API server
func (s *todoAPIServer) Run() error {
	log.Infof("%s %s serving on port %d", s.Name(), common.GetVersion(), s.port)
	return s.engine.Start(fmt.Sprintf(":%d", s.port))
}

// Shutdown stops the Todo API server
func (s *todoAPIServer) Shutdown(ctx context.Context) error {
	log.Infof("shuting down %s %s serving on port %d", s.Name(), common.GetVersion(), s.port)
	return s.engine.Shutdown(ctx)
}
