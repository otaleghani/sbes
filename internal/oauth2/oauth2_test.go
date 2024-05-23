package oauth2

import (
	"testing"
  "fmt"
)

func Test_Oauth(t *testing.T) {
  result := GetOauth2("","")
  fmt.Println(result)
}

