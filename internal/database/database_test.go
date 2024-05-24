package database

import (
  "testing"
  "github.com/otaleghani/spg"
)

var gen = spg.New("en-usa")

func Test_Accounts(t *testing.T) {
  username := gen.Internet().Email(spg.Options{Format: "lower"})

  err := AccountAdd(
    username,
    gen.Internet().Email(spg.Options{Format: "lower"}),
    "smtp." + gen.Internet().DomainName(spg.Options{Format: "lower"}),
    123,
  )
  if err != nil {
    t.Fatal(err)
  }

  err = AccountDelete(username)
  if err != nil {
    t.Fatal(err)
  }
}

func Test_OAuth_tokens(t *testing.T) {
  username := "email@testing.com"
  token := "token..."

  if err := OauthTokenAdd(username, token); err != nil {
    t.Fatal(err)
  }

  if err := OauthTokenDelete(username); err != nil {
    t.Fatal(err)
  }
}

func Test_OAuth_clients(t *testing.T) {
  name := "email@testing.com"
  client_id := "client..."
  client_secret := "secret..."

  if err := OauthClientAdd(name, client_id, client_secret); err != nil {
    t.Fatal(err)
  }

  if err := OauthClientDelete(name); err != nil {
    t.Fatal(err)
  }
}

func Test_Recipients(t *testing.T) {
  name := "clients"
  list := []string{"some@email.com","some@client.com"}

  if err := RecipientsAdd(name, list); err != nil {
    t.Fatal(err)
  }

  if err := RecipientsDelete(name); err != nil {
    t.Fatal(err)
  }
}

func Test_Messages(t *testing.T) {
  helo := "helo"
  body := []byte(helo)
  subject := "subject"
  name := "name"

  // The idea is to make the program tag the message itself
  msg_type := "html"

  if err := MessageAdd(name, subject, msg_type, body); err != nil {
    t.Fatal(err)
  }

  if err := MessageDelete(name); err != nil {
    t.Fatal(err)
  }
}
