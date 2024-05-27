package oauth2

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetAccessToken(clientID, clientSecret, refreshToken string) (string, error) {
	// Create a token with the refresh token.
	// token := &oauth2.Token{RefreshToken: refreshToken}
	config := &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		Endpoint:     google.Endpoint,
	}
	token := &oauth2.Token{
		RefreshToken: "",
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
