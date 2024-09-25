package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
