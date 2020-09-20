package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"net/http"
	"os"

	"github.com/maliur/cardmarket/pkg/http/rest"
	"github.com/maliur/cardmarket/pkg/listing"
	"github.com/maliur/cardmarket/pkg/oauth"
)

func main() {
	l := hclog.Default()
	l.SetLevel(hclog.Debug)

	config, err := getConfig()
	if err != nil {
		l.Error("Failed to read config", "error", err)
		os.Exit(1)
	}

	ls := listing.NewService(config)
	r := rest.NewRouter(l, ls)

	l.Info("Server is running on: http://localhost:8080")
	l.Error("Failed to start server", http.ListenAndServe(":8080", r))
}

func getConfig() (oauth.Config, error) {
	token, ok := os.LookupEnv("APP_TOKEN")
	if !ok {
		return oauth.Config{}, fmt.Errorf("env variable APP_TOKEN not found")
	}
	secret, ok := os.LookupEnv("APP_SECRET")
	if !ok {
		return oauth.Config{}, fmt.Errorf("env variable APP_SECRET not found")
	}
	accessToken, ok := os.LookupEnv("ACCESS_TOKEN")
	if !ok {
		return oauth.Config{}, fmt.Errorf("env variable ACCESS_TOKEN not found")
	}
	accessSecret, ok := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !ok {
		return oauth.Config{}, fmt.Errorf("env variable ACCESS_TOKEN_SECRET not found")
	}

	config := oauth.Config{
		ConsumerKey:       token,
		ConsumerSecret:    secret,
		AccessToken:       accessToken,
		AccessTokenSecret: accessSecret,
	}

	return config, nil
}
