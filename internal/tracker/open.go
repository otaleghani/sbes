package tracker

import (
  "os"
  "time"
  "net/http"
  "log"
  "image"
  "image/png"

  "github.com/otaleghani/sbes/internal/database"
)

func handleOpenTracker(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("email")
  campaign := r.FormValue("campaign")
  log.Printf("Logged open: %30v opended %30v", email, campaign)

  time := time.Now()
  err := database.TrackedOpenAdd(campaign, email, time)
  if err != nil {
    log.Printf("ERROR: unable to save to database entry click of %v in campaign %v", email, campaign)
  }

  http.ServeFile(w, r, "image.png")
}

func createTrackerImage() {
  img := image.NewRGBA(image.Rect(1, 1, 0, 0))

  // Save the image to a file
  file, err := os.Create("image.png")
  if err != nil {
      log.Fatalf("failed to create file: %v", err)
  }
  defer file.Close()

  err = png.Encode(file, img)
  if err != nil {
      log.Fatalf("failed to encode image: %v", err)
  }

  log.Println("Empty image created successfully")
}
