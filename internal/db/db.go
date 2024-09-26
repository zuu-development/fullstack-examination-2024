// Package db provides the database connection and migration functionality.
package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// New creates a new database connection
func New(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
