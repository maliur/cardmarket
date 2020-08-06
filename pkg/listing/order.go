package listing

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
