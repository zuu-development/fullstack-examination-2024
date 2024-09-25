package cmd

import (
	"context"

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
		Run: func(c *cobra.Command, args []string) {
			ctx := c.Context()
			opts := server.TodoServerOpts{
				ListenPort: cfg.Server.Port,
				Config:     cfg,
			}

			server := server.NewServer(ctx, opts)
			for {
				var closer func()
				ctx, cancel := context.WithCancel(ctx)
				server.Run(ctx)
				cancel()
				if closer != nil {
					closer()
				}
			}
		},
	}
	return &serverCmd
}
