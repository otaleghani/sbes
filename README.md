# sbes
Simple bulk email sender

## Features

### send -> sends the message

-r recipient.csv    -> path to recipient list
-a auth             -> setted auth
-m message          -> either txt or html

### auth -> adds auth information

-e email@gmail.com  -> email address
-p password         -> password

## Data storage

``` javascript
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
