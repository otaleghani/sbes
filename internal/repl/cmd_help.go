package repl

import (
	"fmt"
)

func cmdHelp() {
	fmt.Println(`sbes - simple bulk email sender

Usage:  
  sbes [command] [args]

sbes is a tool for sending bulk email directly from the cmdline. Add
your email account, authenticate either with password or OAuth2 and
start sending emails to your mailing lists.

To start sending emails you will need to add an account, a new mailing
list and a message. Optionally you can even add an id and a secret for
OAuth2 authentication. sbes saves every single data locally under
$HOME/.cache/sbes/db.json, so you could even make this program
portable.

Learn more about sbes and its commands by using sbes help with a
specific command, like sbes help add

Commands:
  sbes add [args]       Adds a new record to the local db
  sbes list [args]      Lists the given record
  sbes delete [args]    Deletes given record
  sbes send [args]      Sends email
  `)
}

func cmdHelpAdd() {
	fmt.Println(`sbes add [args]

This command will prompt you to add either a new account, OAuth
client, mailing list or message. In the case of mailing list you will
be asked to give the absolute path to a .csv file containing the
emails. While the message will ask for either a .txt or a .html file.

Args: 
  account               Adds a new account
  oauth-client          Adds a new OAuth client
  mailing-list          Adds a new mailing list, asks for a .csv 
  message               Adds a new message, asks for a .txt or .html
  `)
}

func cmdHelpDelete() {
	fmt.Println(`sbes delete [args]

Propts you to delete a specified database field based on its key. 

Args: 
  account               Deletes a account
  oauth-client          Deletes an OAuth client
  mailing-list          Deletes a mailing list, asks for a .csv 
  message               Deletes a message, asks for a .txt or .html
  `)
}

func cmdHelpList() {
	fmt.Println(`sbes list [args]

Lists all of the items for the given table in the database.

Args: 
  account               Lists every account
  oauth-client          Lists every OAuth client
  mailing-list          Lists every mailing list, asks for a .csv 
  message               Lists every message, asks for a .txt or .html
  `)
}

func cmdHelpSend() {
	fmt.Println(`sbes send [args]

Promps you to specify an account, a mailing list and a message. It
then sends the given email to the mailing list.

Args: 
  password              Sends emails with password authentication
  oauth                 Sends with OAuth authentication
  `)
}
