// Package model provides the data models for the application.
package model

// Config is the configuration for the application.
type Config struct {
	UI     UI
	Server Server
	SQLite SQLite
}

// UI is the configuration for the UI.
type UI struct {
	URL string
}

// Server is the configuration for the server.
type Server struct {
	Port int
}

// SQLite is the configuration for the SQLite database.
type SQLite struct {
	DBFilename string
}
