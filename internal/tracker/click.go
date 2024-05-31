package tracker

import (
  "net/http"
  "log"
  "regexp"
  "fmt"
  "time"

  "github.com/otaleghani/sbes/internal/database"
)

func handleClickTracker(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("email")
  campaign := r.FormValue("campaign")
  redirect := r.FormValue("redirect")
  log.Printf("Logged click: %30v clicked %30v on %30v", email, redirect, campaign)

  time := time.Now()
  err := database.TrackedClickAdd(campaign, email, redirect, time)
  if err != nil {
    log.Printf("ERROR: unable to save to database entry click of %v in campaign %v", email, campaign)
  }

  http.Redirect(w, r, redirect, http.StatusSeeOther)
}

func AddTrackerToEmail(domain, email, campaign, message string) string {
	clicks := regexp.MustCompile(`<a href="([^"]+)">`)
	replacementClicks := fmt.Sprintf(
    `<a href="%s/track/click?email=%s&campaign=%s&redirect=$1">`,
    domain, email, campaign)
	result := clicks.ReplaceAllString(message, replacementClicks)

  open := regexp.MustCompile(`</body>`) 
  replacementOpen := fmt.Sprintf(
    `<img src="%s/track/open?email=%s&campaign=%s"></html>`,
    domain, email, campaign)
  result = open.ReplaceAllString(result, replacementOpen)
	return result
}
