# sbes
Simple bulk email sender

## Commands

### send 

Takes a list of recipient (csv) a email address and a message (either txt or html format) and sends the message to every recipient.

-r recipient.csv    -> path to recipient list
-a auth             -> setted auth
-m message          -> either txt or html

### auth 

Adds a new authorized email address.

-e email@gmail.com  -> email address
-p password         -> password

## To do

- [ ] Data storage in json
- [ ] TestSMTPConnection function
- [ ] Some checks before sending, like is the address duplicate?
- [ ] Test this with other email
- [ ] Encrypt sensitive data

### Data storage

``` javascript
{
    "Accounts": {
        "username": {
            "smtpPort": 587,
            "smtpHost": "smtp.gmail.com",
            "username": "",
            "password": "",
        },
    }
}

{
   "accounts":[
      {
         "username":"",
         "password":""
      },
      {
         "username":"",
         "password":""
      },
   ]
}
```

### TestSMTPConnection

``` go
func TestSMTPConnection(smtpHost string, smtpPort int, username, password string) error {
    d := gomail.NewDialer(smtpHost, smtpPort, username, password)

    // Try to establish a connection
    c, err := d.Dial()
    if err != nil {
        return err
    }
    defer c.Close()
    return nil
}
```
