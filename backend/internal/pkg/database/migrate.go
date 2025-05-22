package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jmoiron/sqlx"
)

// MigrateDB applies all migrations from the migrations directory
func MigrateDB(db *sqlx.DB) error {
	migrationsDir := "migrations"

	// Check if migrations directory exists
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		migrationsDir = filepath.Join("..", migrationsDir)
		if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
			return fmt.Errorf("migrations directory not found")
		}
	}

	// Get all SQL files in the migrations directory
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	// Filter for .sql files and sort by name
	var migrations []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			migrations = append(migrations, file.Name())
		}
	}
	sort.Strings(migrations)

	// Apply migrations in order
	for _, migration := range migrations {
		log.Printf("Applying migration: %s", migration)
		
		// Read migration file
		path := filepath.Join(migrationsDir, migration)
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading migration %s: %w", migration, err)
		}

		// Execute migration
		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("error executing migration %s: %w", migration, err)
		}
		
		log.Printf("Successfully applied migration: %s", migration)
	}

	return nil
}
