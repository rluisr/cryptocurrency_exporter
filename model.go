package main

type Request struct {
	Currency string `json:"currency"`
	Code     string `json:"code"`
	Meta     bool   `json:"meta"`
}

type Response struct {
	Rate      float64 `json:"rate"`
	Volume    float64 `json:"volume"`
	Cap       float64 `json:"cap"`
	Liquidity float64 `json:"liquidity"`
	Delta     struct {
		Hour    interface{} `json:"hour"`
		Day     interface{} `json:"day"`
		Week    interface{} `json:"week"`
		Month   interface{} `json:"month"`
		Quarter interface{} `json:"quarter"`
		Year    interface{} `json:"year"`
	} `json:"delta"`
}
