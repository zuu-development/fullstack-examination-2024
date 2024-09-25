package server

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/zuu-development/fullstack-examination-2024/internal/common"
	"github.com/zuu-development/fullstack-examination-2024/internal/db"
	"github.com/zuu-development/fullstack-examination-2024/internal/handler"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
)

// TodoServer is the API server for Todo
type TodoServer struct {
	port   int
	engine *echo.Echo
	log    *log.Entry
	db     *gorm.DB

	// stopCh is the channel which when closed, will shutdown the Todo server
	stopCh chan struct{}
}

type TodoServerOpts struct {
	ListenPort int
	Config     model.Config
}

// NewServer returns a new instance of the Todo API server
func NewServer(ctx context.Context, opts TodoServerOpts) *TodoServer {
	logger := log.NewEntry(log.StandardLogger())

	dbInstance, err := db.New(opts.Config.SQLite.DBFilename)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	s := &TodoServer{
		port:   opts.ListenPort,
		engine: echo.New(),
		log:    logger,
		db:     dbInstance,
	}

	handler.Register(s.engine, s.db)
	return s
}

// TODO ctx is not used
func (s *TodoServer) Run(ctx context.Context) {

	go func() { s.checkServeErr("api", s.engine.Start(fmt.Sprintf(":%d", s.port))) }()

	// Start the muxed listeners for our servers
	log.Infof("todo %s serving on port %d",
		common.GetVersion(), s.port)

	s.stopCh = make(chan struct{})
	<-s.stopCh

}

// checkServeErr checks the error from a .Serve() call to decide if it was a graceful shutdown
func (s *TodoServer) checkServeErr(name string, err error) {
	if err != nil {
		if s.stopCh == nil {
			// a nil stopCh indicates a graceful shutdown
			log.Infof("graceful shutdown %s: %v", name, err)
		} else {
			log.Fatalf("%s: %v", name, err)
		}
	} else {
		log.Infof("graceful shutdown %s", name)
	}
}
