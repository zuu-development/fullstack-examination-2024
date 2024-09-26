// Package model provides the data models for the application.
package model

// Config is the configuration for the application.
type Config struct {
	Server Server
	SQLite SQLite
}

// Server is the configuration for the server.
type Server struct {
	Port int
}

// SQLite is the configuration for the SQLite database.
type SQLite struct {
	DBFilename string
}
