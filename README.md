# Simple Bulk Email Sender
[![Go Report Card](https://goreportcard.com/badge/github.com/otaleghani/sbes)](https://goreportcard.com/report/github.com/otaleghani/sbes)
![Ci tests](https://github.com/otaleghani/sbes/actions/workflows/tests.yml/badge.svg)

Simple bulk email sender. A command line utility to send bulk email to your mailing lists right in your terminal.

## Features

- Multiple accounts: Save access of multiple accounts in one place
- OAuth: Connect with the GC API to send emails with your Gmail account
- Easy to use: Every single command guide you through the configuration
- Portable: Everything sits inside a single json file

## Installation

You will need Go at least 1.2x. Afterwards you can `go install` sbes.

```
go install github.com/otaleghani/sbes
```

## Basic usage

Every single command waits for your input, so you don't have to learn a bunch of flags.

To get started, add a new account with `add account`. It will ask for your username, password, SMTP host and port. Your data is store locally and never leaves your pc.

``` bash
sbes add account
```

You can then add a new OAuth2 token for your account. This is optional, but it's required with some email providers, like Gmail. You will need to generate from Google Cloud a new project and get the OAuth2 Client Id and Secret. You can find a more in-depth explaination at [How to setup OAuth2](#how-to-setup-oauth2). Currently only GMail API is supported.

```  bash
sbes add oauth-client       # Adds a new OAuth2 client
sbes add oauth-token        # Adds token to an existing account
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

To use OAuth2 with gomail.v2, you'll need to:

1. Set up a Google Cloud Project and enable the Gmail API.
2. Obtain OAuth2 credentials.
3. Add the OAuth Client to sbes

### 1. Set up a Google Cloud Project and enable the Gmail API

1. Go to the Google Cloud Console.
2. Create a new project.
3. Enable the Gmail API for the project.
4. Set up the OAuth consent screen.

### 2. Obtain OAuth2 credentials

1. In the Google Cloud Console, go to Credentials.
2. Create OAuth 2.0 Client IDs. Choose "Web application" as the application type.
3. Add redirect URI, that would be `http://localhost:8080` and `http://localhost:8080/oauth2callback`
4. Save this new OAuth Client and the Id and Secret should pop out.

### 3. Add the OAuth Client to sbes

1. Go back to your terminal
2. Create a new OAuth Client into sbes (`sbes add oauth-client`) using your newly created Id and Secret

You can now generate a new token for an account by using `sbes add oauth-token`. You will be promped with selecting the account that you want to authenticate with OAuth and a OAuth Client. Afterwars you will need to visit localhost:8080 and follow up with the OAuth consent screen.

This first run adds both a Refresh token and an Access token. The Access token is valid for 1 hour and does not refresh automatically (as of now). To refresh it you will need to use the cmd `sbes refresh`. After you indicate what account you want to refresh, it will update the access token with a new one.

When the Refresh token runs out (every 7 days), you will need to call the `sbes add oauth-token` again.

## Availables commands

- `add` - Adds either a new `account`, `mailing-list`, `message`, `oauth-client` or `oauth-token`

- `delete` - Deletes either an `account`, `mailing-list`, `message`, or `oauth-client` 

- `list` - Lists either an `account`, `mailing-list`, `message`, or `oauth-client` 

- `send` - You will need to specify either `password` or `oauth` for the authentication method.

## Planning for future versions

### 0.1.1
    tdb

### Features to add
- [ ] Encrypt local sensitive data
- [ ] Add other APIs, like Microsoft Graph API
- [ ] Trackers 

#### Trackers brainstorming
Requires a running process that serves a resource like an image. Could do something like `sbes send...` ask then for a name and then create with that name a hittable resource (like an image) and then make it available to the recipients. This image is then inserted inside of every html email with a parameter (like the email).

`somehost.com/trackers/[campaign-name].png`
e.g. tracker

`somehost.com/trackers/[campaign-name].png?e=some@email.com`
e.g. full tracker

Then whenever that image is hit it has this arg that contain the email address of the reciever. You could then parse it and add it inside of a log file with that campaign name.

``` bash
sbes track
```
