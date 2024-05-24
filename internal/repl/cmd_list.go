package repl

import (
  "github.com/otaleghani/sbes/internal/database"
  "fmt"
)

func cmdList_Accounts() {
  if err := database.AccountsList(); err != nil {
    fmt.Println("ERROR: ", err)
    return
  }
}

func cmdList_OAuthClients() {
  if err := database.OAuthClientsList(); err != nil {
    fmt.Println("ERROR: ", err)
    return
  }
}

func cmdList_OAuthTokens() {
  if err := database.OAuthTokensList(); err != nil {
    fmt.Println("ERROR: ", err)
    return
  }
}

func cmdList_MailingLists() {
  if err := database.MailingListsList(); err != nil {
    fmt.Println("ERROR: ", err)
    return
  }
}

func cmdList_Messages() {
  if err := database.MessagesList(); err != nil {
    fmt.Println("ERROR: ", err)
    return
  }
}
