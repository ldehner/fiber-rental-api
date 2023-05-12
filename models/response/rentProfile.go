package responsemodels

type RentProfile struct {
	ContractDue string  `json:"ContractDue"`
	HeatCosts   float32 `json:"HeatCosts"`
	Deposit     float32 `json:"Deposit"`
	Id          string  `json:"Id"`
	Minimum     uint8   `json:"Minimum"`
	Rent        float32 `json:"Rent"`
	RentalStart string  `json:"RentalStart"`
}
