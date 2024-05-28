// Helper function to read either a .html or .txt file and parse it's
// content in a string

package parser

import (
	"errors"
	"os"
	"path/filepath"
)

func Email(path string) (string, string, error) {
	// Clean the input path
	cleanPath := filepath.Clean(path)

	baseDir := "/your/base/directory" // Change to your base directory
	if !filepath.IsAbs(cleanPath) {
		cleanPath = filepath.Join(baseDir, cleanPath)
	}

	// Takes extension of file in given path
	extension := filepath.Ext(cleanPath)

	// Checks if .html or .txt
	if extension != ".html" && extension != ".txt" {
		return "", "", errors.New("the file must be .html or .txt")
	}

	// Reads file
	file, err := os.ReadFile(cleanPath)
	if err != nil {
		return "", "", err
	}

	// Returns string
	return string(file), extension, nil
}
