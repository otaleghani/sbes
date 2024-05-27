package oauth2

import (
	"fmt"
	"testing"
)

func Test_Oauth(t *testing.T) {
	//result := GetOauth2("", "")
	//fmt.Println(result)
	result, _ := GetAccessToken("", "", "thetoken")

	fmt.Println(result)
}
