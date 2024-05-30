package tracker

import (
  "net/http"
  "log"
  "regexp"
  "fmt"
)

func handleClickTracker(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("email")
  campaign := r.FormValue("campaign")
  redirect := r.FormValue("redirect")
  log.Printf("Logged click: %30v clicked %30v on %30v", email, redirect, campaign)

  // Here save result if ok

  http.Redirect(w, r, redirect, http.StatusSeeOther)
}

func AddTrackerToEmail(domain, email, campaign, message string) string {
	re := regexp.MustCompile(`<a href="([^"]+)">`)
	replacement := fmt.Sprintf(
    `<a href="%s/track/click?email=%s&campaign=%s&redirect=$1">`,
    domain, email, campaign)
	result := re.ReplaceAllString(message, replacement)
	return result
}
