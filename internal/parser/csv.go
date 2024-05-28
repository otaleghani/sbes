// Helper functions used to parse emails from csv

package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
  "path/filepath"
)

func Csv(path string, columnNumber int) ([]string, error) {
  // Clean the input path
  cleanPath := filepath.Clean(path)

  baseDir := "/your/base/directory"  // Change to your base directory
  if !filepath.IsAbs(cleanPath) {
    cleanPath = filepath.Join(baseDir, cleanPath)
  }

	// Opens file at given path
	file, err := os.Open(cleanPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Creates a new csv Reader
	reader := csv.NewReader(file)
	var columnData []string

	// Cycles through every row
	for {
		// Reads the data of the row
		record, err := reader.Read()
		if err != nil {
			// If finds end of file, break
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		// If the given column number is negative or out or reach, returns
		if columnNumber < 0 || columnNumber >= len(record) {
			return nil, fmt.Errorf("column number %d out of range", columnNumber)
		}

		// If the given cell doesn't have a "@" sign, ignore it
		if !strings.Contains(record[columnNumber], "@") {
			continue
		}

		// Appends cell data
		columnData = append(columnData, record[columnNumber])
	}

	// Return []string filled with data
	return columnData, nil
}
