package model

type Config struct {
	Server Server
	SQLite SQLite
}
type Server struct {
	Port int
}

type SQLite struct {
	DBFilename string
}
