package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func SetupConfig() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     "18988669711-nas9to1c6oba181d8qk2po389ol0ikoi.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:8080/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}
