package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/zuu-development/fullstack-examination-2024/internal/common"
)

// swaggerServer is the API server for Todo
type swaggerServer struct {
	port   int
	engine *echo.Echo
	log    *log.Entry
}

// SwaggerServerOpts is the options for the swaggerServer
type SwaggerServerOpts struct {
	ListenPort int
}

// NewSwagger returns a new instance of the Swagger server
func NewSwagger(opts SwaggerServerOpts) Server {
	logger := log.NewEntry(log.StandardLogger())
	log.SetFormatter(&log.JSONFormatter{})

	engine := echo.New()
	engine.HideBanner = true
	engine.HidePort = true

	engine.Use(requestLogger())

	engine.GET("/swagger/*", echoSwagger.WrapHandler)

	s := &swaggerServer{
		port:   opts.ListenPort,
		engine: engine,
		log:    logger,
	}

	return s
}

func (s *swaggerServer) Name() string {
	return "swaggerServer"
}

func (s *swaggerServer) Run() error {
	log.Infof("%s %s serving on port %d", s.Name(), common.GetVersion(), s.port)
	return s.engine.Start(fmt.Sprintf(":%d", s.port))
}

func (s *swaggerServer) Shutdown(ctx context.Context) error {
	log.Infof("shuting down %s %s serving on port %d", s.Name(), common.GetVersion(), s.port)
	return s.engine.Shutdown(ctx)
}
