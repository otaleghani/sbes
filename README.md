# sbes
[![Go Report Card](https://goreportcard.com/badge/github.com/otaleghani/sbes)](https://goreportcard.com/report/github.com/otaleghani/sbes)
![Ci tests](https://github.com/otaleghani/sbes/actions/workflows/tests.yml/badge.svg)

Simple bulk email sender. A command line utility to send bulk email to your mailing lists right in your terminal.

## Basic usage

Every single command waits for your input, so you don't have to learn a bunch of flags.

To get started, add a new account with `add account`. It will ask for your username, password, SMTP host and port. Your data is store locally and never leaves your pc.

``` bash
sbes add account
```

You can then add a new OAuth2 token for your account. This is optional, but it's required with some email providers, like Gmail. You will need to generate from Google Cloud a new project and get the OAuth2 Client Id and Secret. You can find a more in-depth explaination at ![How to setup OAuth2](#how-to-setup-oauth2).

```  bash
sbes add oauth-client
sbes add oauth-token
```

Afterwards you will need to specify a mailing list and a message to send. Here you will need a `.csv` file for the mailing list and either a `.txt` or `.html` file for the message. Provide to sbes the absolute path.

Every single data is saved on a `.json` file, so you can even make sbes portable.

```  bash
sbes add mailing-list
sbes add message
```

Finally you can specify to sbes which message to send from which account to which mailing list.

```  bash
sbes send password
sbes send oauth
```

## How to setup OAuth2

1. Create a Google Cloud account
2. New project
3. Activate Gmail API
4. Create a new OAuth2 client for this app

## Commands

### add 

### delete

### list

### send

## To do

Before 0.1.0

- [ ] Validate input before adding to database
- [ ] Create a better print format fro add, delete and send
- [ ] Encrypt sensitive data
- [ ] README.md
- [ ] Wiki

### Data storage

