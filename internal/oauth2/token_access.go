// Function to get the access token from a specific refresh token

package oauth2

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetAccessToken(clientID, clientSecret, refreshToken string) (string, error) {
	// Sets the oauth2.Config
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     google.Endpoint,
	}

	// Initializes the variable that holds the refresh token
	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// Create a token source with the refresh tokene
	tokenSource := config.TokenSource(context.Background(), token)

	// Retrieve a new access token.
	newToken, err := tokenSource.Token()
	if err != nil {
		return "", err
	}

	// Returns the access token
	return newToken.AccessToken, nil
}
