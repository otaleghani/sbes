// Functions used to get a refresh token from Google OAuth2

package oauth2

import (
	"context"
	"fmt"
	"log"
	"net/http"
  "time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Define the OAuth2 configuration. This includes the client ID,
// client secret, redirect URL, the scopes of access required, and the
// OAuth2 provider's endpoint.
var (
	oauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8080/oauth2callback",
		Scopes:      []string{"https://mail.google.com/"},
		Endpoint:    google.Endpoint,
	}
	// State token to protect against CSRF
	oauthStateString = "state-token"

	// Channel to receive OAuth2 tokens
	tokenChan = make(chan *oauth2.Token)
)

// GetOauth initializes the OAuth2 configuration with provided client
// ID and secret, sets up HTTP handlers, and starts the web server. It
// waits for the token to be received through the token channel and
// returns the access token.

func GetOauth2(id, secret string) (string, string) {
	oauthConfig.ClientID = id
	oauthConfig.ClientSecret = secret

	// Set up HTTP handlers for the root and the OAuth2 callback
	// endpoint.
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/oauth2callback", handleOAuth2Callback)

	// Start the web server in a new goroutine.
  srv := &http.Server{
      Addr: ":8080",
      Handler: nil,
      ReadTimeout: 5 * time.Second,
      WriteTimeout: 10 * time.Second,
      IdleTimeout: 15 * time.Second,
  }
	go func() {
		fmt.Println("Visit the URL for the auth dialog: http://localhost:8080")
    log.Fatal(srv.ListenAndServe())
		//log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Wait for the token to be received on the token channel.
	token := <-tokenChan

	fmt.Println(token.RefreshToken)

	return token.RefreshToken, token.AccessToken
}

// handleMain redirects the user to the Google OAuth2 authorization
// URL.
func handleMain(w http.ResponseWriter, r *http.Request) {
	// Generate the URL for the OAuth2 authorization request.
	url := oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	// Redirect the user to the authorization URL.
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// handleOAuth2Callback handles the OAuth2 callback from the provider.
// It validates the state parameter, exchanges the authorization code
// for an access token, and sends the token through the token channel.
func handleOAuth2Callback(w http.ResponseWriter, r *http.Request) {
	// Verify the state parameter to protect against CSRF.
	state := r.FormValue("state")
	if state != oauthStateString {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Error(w, "Invalid state parameter", http.StatusBadRequest)
		return
	}

	// Get the authorization code from the request.
	code := r.FormValue("code")
	// Exchange the authorization code for an access token.
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("oauthConfig.Exchange() failed with '%s'\n", err)
		http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
		return
	}

	// Print the access token and send it through the token channel.
	fmt.Fprintf(w, "Access token: %s\n", token.AccessToken)
	tokenChan <- token
}
