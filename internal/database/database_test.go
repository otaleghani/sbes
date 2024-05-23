package database

import (
  "testing"
  "github.com/otaleghani/spg"
)

var gen = spg.New("en-usa")

func Test_Auth(t *testing.T) {
  username := gen.Internet().Email(spg.Options{Format: "lower"})

  err := AuthAdd(
    username,
    gen.Internet().Email(spg.Options{Format: "lower"}),
    "smtp." + gen.Internet().DomainName(spg.Options{Format: "lower"}),
    123,
  )
  if err != nil {
    t.Fatal(err)
  }

  err = AuthDelete(username)
  if err != nil {
    t.Fatal(err)
  }
}

