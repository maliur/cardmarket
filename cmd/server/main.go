package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maliur/cardmarket/pkg/http/rest"
	"github.com/maliur/cardmarket/pkg/listing"
	"github.com/maliur/cardmarket/pkg/oauth"
)

func main() {
	var config oauth.Config
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	err = json.NewDecoder(jsonFile).Decode(&config)
	if err != nil {
		log.Fatalf("could not unmarshal config: %v", err)
	}

	l := listing.NewService(config)
	router := rest.NewRouter(l)

	log.Println("Server is running on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
