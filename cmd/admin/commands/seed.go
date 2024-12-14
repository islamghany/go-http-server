package commands

import (
	"database/sql"
	"fmt"
	"httpserver/internal/config"
	"os"
	"path/filepath"
)

func SeedDB(cfg *config.Config, dbConn *sql.DB) error {
	seedDir := "internal/db/seeds"

	// Get list of seed files
	files, err := getSeedFiles(seedDir)
	if err != nil {
		return fmt.Errorf("failed to get seed files: %w", err)
	}

	// Execute each seed file
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read seed file %s: %w", file, err)
		}
		_, err = dbConn.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to execute seed file %s: %w", file, err)
		}
	}

	return nil
}

func getSeedFiles(seedDir string) ([]string, error) {
	var files []string

	// Read the directory
	dir, err := os.ReadDir(seedDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read seed directory: %w", err)
	}

	// Get all the files
	for _, file := range dir {
		if file.IsDir() {
			continue
		}
		files = append(files, filepath.Join(seedDir, file.Name()))
	}

	return files, nil
}
