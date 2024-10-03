// Package cmd provides the command line interface for the application.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/zuu-development/fullstack-examination-2024/internal/db"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Run: func(_ *cobra.Command, _ []string) {
		dbDir := filepath.Dir(cfg.SQLite.DBFilename)
		if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
			log.Fatalf("failed to create directory: %s err: %s", dbDir, err)
			return
		}

		dbInstance, err := db.New(cfg.SQLite.DBFilename)
		if err != nil {
			log.Fatalf("failed to open database filename: %s err: %s", cfg.SQLite.DBFilename, err)
			return
		}
		if err := db.Migrate(dbInstance); err != nil {
			log.Fatalf("failed to migrate database err: %s", err)
		}
		fmt.Println("Migration completed. SQLite.DBFilename: ", cfg.SQLite.DBFilename)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
