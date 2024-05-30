package database

import (
	// "github.com/otaleghani/spg"
	"fmt"
	"testing"
  "time"
)

// var gen = spg.New("en-usa")

func Test_Paths(t *testing.T) {
	path, fileName, err := databasePath()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(path, fileName)
}

func Test_Campaigns(t *testing.T) {
  // err := CampaignAdd("sandro3", "o.taleghani@talesign.com", "test", "test")
	// if err != nil {
	// 	t.Fatal(err)
	// }

  err := CampaignDelete("sandro3")
	if err != nil {
		t.Fatal(err)
	}

  // time := time.Now()
  // err = TrackedOpenAdd("sandro", "sandro@fortissimo.com", time)
	// if err != nil {
	// 	t.Fatal(err)
	// }

  time := time.Now()
  err = TrackedClickAdd("sandro", "sandro@fortissimo.com", "linkpotente2", time)
	if err != nil {
		t.Fatal(err)
	}

  err = CampaignList()
	if err != nil {
		t.Fatal(err)
	}
}
