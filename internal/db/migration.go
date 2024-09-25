package db

import (
	"github.com/zuu-development/fullstack-examination-2024/internal/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.Todo{}); err != nil {
		return err
	}
	return nil
}
