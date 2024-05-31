package tracker

import (
	"log"
	"testing"
)

func Test_AddTracker(t *testing.T) {
	s := AddTrackerToEmail("localhost:8081", "o.taleghani@talesign.com", "nomecampagna", "some incredibbile testo <a href=\"https://aficleaning.com/\"></a>")
	log.Println(s)
}

func Test_Tracker(t *testing.T) {
	Listen()
}
