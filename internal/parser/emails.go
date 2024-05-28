// Helper function to read either a .html or .txt file and parse it's
// content in a string

package parser

import (
	"errors"
	"os"
	"path/filepath"
)

func Email(path string) (string, string, error) {
  // Takes extension of file in given path
	extension := filepath.Ext(path)

  // Checks if .html or .txt
	if extension != ".html" && extension != ".txt" {
		return "", "", errors.New("The file must be .html or .txt")
	}

  // Reads file
	file, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}

  // Returns string
	return string(file), extension, nil
}
