package tracker

import (
  "net/http"
  "log"
  "time"
)

func Listen() {
  createTrackerImage()
  http.HandleFunc("/track/open", handleOpenTracker)
  http.HandleFunc("/track/click", handleClickTracker)

  srv := &http.Server{
    Addr: ":8081",
    Handler: nil,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout: 15 * time.Second,
  }

  for {
    log.Println("Serving tracker on port :8081")
    err := srv.ListenAndServe()
    if err != nil {
      log.Println(err)
    }
  }
}
