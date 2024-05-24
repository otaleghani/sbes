package parser

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func Email(path string) (string, string, error) {
	extension := filepath.Ext(path)
	fmt.Println(extension)

	if extension != ".html" && extension != ".txt" {
		return "", "", errors.New("The file must be .html or .txt")
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return "", "", err
	}
	return string(file), extension, nil
}
