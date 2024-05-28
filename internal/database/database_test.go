package database

import (
	// "github.com/otaleghani/spg"
	"testing"
  "fmt"
)

// var gen = spg.New("en-usa")

func Test_Paths(t *testing.T) {
  path, fileName, err := databasePath()
  if err != nil {
    t.Fatal(err)
  }
  fmt.Println(path, fileName)
}
