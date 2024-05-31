package repl

import (
	"fmt"

	"github.com/otaleghani/sbes/internal/database"
	"github.com/otaleghani/sbes/internal/tracker"
)

func cmdTrackerStart(domain string) {
	err := database.UpdateDomainTracker(domain)
	if err != nil {
		fmt.Println("ERROR: unable to add the domain")
		return
	}

	tracker.Listen()
}
