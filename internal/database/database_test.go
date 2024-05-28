package database

import (
	// "github.com/otaleghani/spg"
	"fmt"
	"testing"
)

// var gen = spg.New("en-usa")

func Test_Paths(t *testing.T) {
	path, fileName, err := databasePath()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(path, fileName)
}
