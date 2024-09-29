package server

import (
	"testing"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"gorm.io/gorm"
)

func TestNewAPI(t *testing.T) {
	tests := []struct {
		name    string
		opts    TodoAPIServerOpts
		wantErr bool
	}{
		{
			name: "Valid configuration",
			opts: TodoAPIServerOpts{
				ListenPort: 8080,
				Config: model.Config{
					SQLite: model.SQLite{
						DBFilename: ":memory:",
					},
					UI: model.UI{
						URL: "http://localhost:3000",
					},
					SwaggerServer: model.Server{
						Enable: true,
						Port:   8081,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid database configuration",
			opts: TodoAPIServerOpts{
				ListenPort: 8080,
				Config: model.Config{
					SQLite: model.SQLite{
						DBFilename: "/invalid/path/to/db",
					},
					UI: model.UI{
						URL: "http://localhost:3000",
					},
					SwaggerServer: model.Server{
						Enable: true,
						Port:   8081,
					},
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, err := NewAPI(tt.opts)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, server)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, server)
				assert.Equal(t, tt.opts.ListenPort, server.(*todoAPIServer).port)
				assert.IsType(t, &echo.Echo{}, server.(*todoAPIServer).engine)
				assert.IsType(t, &log.Entry{}, server.(*todoAPIServer).log)
				assert.IsType(t, &gorm.DB{}, server.(*todoAPIServer).db)
			}
		})
	}
}
