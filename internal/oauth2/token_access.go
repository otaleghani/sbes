package oauth2

import (
  "fmt"
  "context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetAccessToken(clientID, clientSecret, refreshToken string) (string, error) {
	// Create a token with the refresh token.
	// token := &oauth2.Token{RefreshToken: refreshToken}
  config := &oauth2.Config{
    ClientID:     "1029956480466-f2p5kkp4dsahbtrh0q1svm0mvuu3dsja.apps.googleusercontent.com",
		ClientSecret: "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",
		Endpoint:    google.Endpoint,
  }
  token := &oauth2.Token{
    RefreshToken: "1//09-lyTNr5WI27CgYIARAAGAkSNwF-L9IrOme_pyxAHCPTEMriKYxg4YJ5141xQHmI9fnT3TPg300rYBzhb5V0d8L-ohwTZfFz1wk",
  }

	// Create a token source with the refresh token.
	tokenSource := config.TokenSource(context.Background(), token)

	// Retrieve a new access token.
	newToken, err := tokenSource.Token()
	if err != nil {
		return "", err
	}

  fmt.Println(token)
	return newToken.AccessToken, nil
}
