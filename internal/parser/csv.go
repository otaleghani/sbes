package parser

import (
	"encoding/csv"
	"os"
  "fmt"
  "strings"
)

func Csv(path string, columnNumber int) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var columnData []string

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		if columnNumber < 0 || columnNumber >= len(record) {
			return nil, fmt.Errorf("column number %d out of range", columnNumber)
		}

    if !strings.Contains(record[columnNumber], "@") {
      continue
    }
		columnData = append(columnData, record[columnNumber])
	}

	return columnData, nil
}
