package models

type Property struct {
	ID       uint    `json:"id"`
	Tenant   User    `json:"tenant_id"`
	Landlord User    `json:"landlord_id"`
	Country  string  `json:"country"`
	City     string  `json:"city"`
	Street   string  `json:"street"`
	Number   string  `json:"number"`
	Zipcode  string  `json:"zipcode"`
	Price    float64 `json:"price"`
	Deposit  float64 `json:"deposit"`
}
