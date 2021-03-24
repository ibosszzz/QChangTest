package model

type CashierDesk struct {
	BankNote1000 int `json:"1000"`
	BankNote500 int `json:"500"`
	BankNote100 int `json:"100"`
	BankNote50 int `json:"50"`
	BankNote20 int `json:"20"`
	Coin10 int `json:"10"`
	Coin5 int `json:"5"`
	Coin1 int `json:"1"`
	Coin025 int `json:"0.25"`
}

type CustomerPay struct {
	ProductPrice float64 `json:"product_price"`
	CustomerPay float64 `json:"customer_pay"`
}