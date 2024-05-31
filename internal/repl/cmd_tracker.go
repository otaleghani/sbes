package repl

import (
  "fmt"

  "github.com/otaleghani/sbes/internal/tracker"
  "github.com/otaleghani/sbes/internal/database"
)

func cmdTrackerStart(domain string) {
  err := database.UpdateDomain(domain)
  if err != nil {
    fmt.Println("ERROR: unable to add the domain")
    return
  }

  tracker.Listen()
}
