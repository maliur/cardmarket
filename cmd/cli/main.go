package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/maliur/cardmarket/oauth"
)

type Article struct {
	Price     float32 `json:"price"`
	Count     int     `json:"count"`
	Condition string  `json:"condition"`
	IsFoil    bool    `json:"isFoil"`
	Product   struct {
		Name      string `json:"enName"`
		Expansion string `json:"expansion"`
		Rarity    string `json:"rarity"`
	} `json:"product"`
}

type OrderState struct {
	State  string `json:"state"`
	Bought string `json:"dateBought"`
	Paid   string `json:"datePaid"`
	Sent   string `json:"dateSent"`
}

type OrderSeller struct {
	Username string `json:"username"`
	Address  struct {
		CountryCode string `json:"country"`
	} `json:"address"`
}

type Order struct {
	IdOrder        int         `json:"idOrder"`
	TrackingNumber string      `json:"trackingNumber"`
	Articles       []Article   `json:"article"`
	State          OrderState  `json:"state"`
	Seller         OrderSeller `json:"seller"`
	ArticleValue   float32     `json:"articleValue"`
	TotalValue     float32     `json:"totalValue"`
}

type Orders struct {
	Orders []Order `json:"order"`
}

func Get(url string, config oauth.Config) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", oauth.OauthHeader(url, config))
	req.Header.Add("Accept", "*/*")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

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

	baseUrl := "https://api.cardmarket.com/ws/v2.0/output.json"

	resp, err := Get(baseUrl+"/orders/buyer/sent", config)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}

	defer resp.Body.Close()

	var orders Orders
	err = json.NewDecoder(resp.Body).Decode(&orders)
	if err != nil {
		log.Fatalf("could not unmarshal response: %v", err)
	}

	max, _ := strconv.Atoi(resp.Header.Get("X-Request-Limit-Max"))
	used, _ := strconv.Atoi(resp.Header.Get("X-Request-Limit-Count"))
	left := max - used
	fmt.Println("Request left today:", left)

	fmt.Println("Total orders:", len(orders.Orders))

	for _, order := range orders.Orders {
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
