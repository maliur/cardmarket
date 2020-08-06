package listing

import (
	"encoding/json"
	"net/http"

	"github.com/maliur/cardmarket/pkg/oauth"
)

// max, _ := strconv.Atoi(resp.Header.Get("X-Request-Limit-Max"))
// used, _ := strconv.Atoi(resp.Header.Get("X-Request-Limit-Count"))
// left := max - used
// fmt.Println("Request left today:", left)

var baseUrl string = "https://api.cardmarket.com/ws/v2.0/output.json"

type Service interface {
	GetSentOrders() ([]Order, error)
	GetPaidOrders() ([]Order, error)
}

// internal service struct to hide eg repository/storage from the client
type service struct {
	c oauth.Config
}

func NewService(c oauth.Config) Service {
	return &service{c}
}

func get(url string, config oauth.Config) (*http.Response, error) {
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

func (s *service) GetSentOrders() ([]Order, error) {
	resp, err := get(baseUrl+"/orders/buyer/sent", s.c)
	if err != nil {
		return nil, err
	}

	var data Orders
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Orders, nil
}

func (s *service) GetPaidOrders() ([]Order, error) {
	resp, err := get(baseUrl+"/orders/buyer/paid", s.c)
	if err != nil {
		return nil, err
	}

	var data Orders
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Orders, nil
}
