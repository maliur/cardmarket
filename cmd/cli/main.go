package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

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

	orders, err := l.GetSentOrders()
	if err != nil {
		log.Fatalf("could not get orders: %v", err)
	}

	for _, order := range orders {
		fmt.Println("----------------------------------")
		fmt.Println("Id:\t\t", order.IdOrder)
		fmt.Println("Seller:\t\t", order.Seller.Username)
		fmt.Println("Country:\t", order.Seller.Address.CountryCode)
		if len(order.TrackingNumber) > 0 {
			fmt.Println("Tracking:\t", order.TrackingNumber)
		}
		fmt.Println("State:\t\t", order.State.State)
		fmt.Println("Bought:\t\t", order.State.Bought)
		fmt.Println("Paid:\t\t", order.State.Paid)
		fmt.Println("Sent:\t\t", order.State.Sent)
		fmt.Printf("Value:\t\t %.2f Euro\n", order.ArticleValue)
		fmt.Printf("Total:\t\t %.2f Euro\n", order.TotalValue)

		fmt.Println("Articles:")
		for _, article := range order.Articles {
			fmt.Println("    Name:", article.Product.Name)
			fmt.Println("    Expansion:", article.Product.Expansion)
			fmt.Println("    Rarity:", article.Product.Rarity)
			fmt.Println("    Count:", article.Count)
			fmt.Printf("    Price: %.2f Euro\n", article.Price)
			fmt.Println("    Condition:", article.Condition)
			fmt.Println("")
		}
	}
}
