package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maliur/cardmarket/pkg/http/rest"
	"github.com/maliur/cardmarket/pkg/listing"
	"github.com/maliur/cardmarket/pkg/oauth"
)

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatalf("Failed to read config: %s", err)
	}

	l := listing.NewService(config)
	router := rest.NewRouter(l)

	log.Println("Server is running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getConfig() (oauth.Config, error) {
	token, present := os.LookupEnv("APP_TOKEN")
	if !present {
		return oauth.Config{}, fmt.Errorf("env variable APP_TOKEN not found")
	}
	secret, present := os.LookupEnv("APP_SECRET")
	if !present {
		return oauth.Config{}, fmt.Errorf("env variable APP_SECRET not found")
	}
	accessToken, present := os.LookupEnv("ACCESS_TOKEN")
	if !present {
		return oauth.Config{}, fmt.Errorf("env variable ACCESS_TOKEN not found")
	}
	accessSecret, present := os.LookupEnv("ACCESS_TOKEN_SECRET")
	if !present {
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
