package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zuu-development/fullstack-examination-2024/internal/server"
)

func init() {
	rootCmd.AddCommand(NewServerCmd())
}

// NewServerCmd returns a new `server` command to be used as a sub-command to root
func NewServerCmd() *cobra.Command {
	serverCmd := cobra.Command{
		Use:   "server",
		Short: "Print version information",
		Example: `  # Print the full version of client and server to stdout
  todo-cli server
`,
		Run: func(_ *cobra.Command, _ []string) {
			var servers []server.Server

			apiOpts := server.TodoAPIServerOpts{
				ListenPort: cfg.APIServer.Port,
				Config:     cfg,
			}
			apiServer, err := server.NewAPI(apiOpts)
			if err != nil {
				log.Fatal(err)
			}
			servers = append(servers, apiServer)

			if cfg.SwaggerServer.Enable {
				SwaggerOpts := server.SwaggerServerOpts{
					ListenPort: cfg.SwaggerServer.Port,
				}
				swagServer := server.NewSwagger(SwaggerOpts)
				servers = append(servers, swagServer)
			}

			ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
			defer stop()

			for _, s := range servers {
				server := s
				go func() {
					if err := server.Run(); err != nil && err != http.ErrServerClosed {
						log.Fatal("shutting down ", server.Name(), " err: ", err)
					}
				}()
			}

			log.Info("server started")
			// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
			<-ctx.Done()
			log.Info("server shutting down")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			for _, s := range servers {
				if err := s.Shutdown(ctx); err != nil {
					log.Fatal(err)
				}
			}
			log.Info("server shutdown gracefully")
		},
	}
	return &serverCmd
}
