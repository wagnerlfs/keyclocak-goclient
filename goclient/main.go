package main

import (
	"context"
	"log"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     = ""
	clientSecret = "7namJR85xEhE3B8cuv81jhImnDsUgGe2"
)

func main() {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "http://localhost:8080/realms/myrealm")

	/*
		if (err != null) {
			log.Fatalf(err)
		}
	*/

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://localhost:8081/auth/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "roles"},
	}

	state := "123"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)

	})

	log.Fatal(http.ListenAndServe("8081", nil))

	/*
		http.HundleFunc("/", func(http.ResponseWriter, request, http.Request){
			 http.Redirect(writer, request, config.AuthCodeURL(state), http.StatusFound)
			}
		)
	*/

}
