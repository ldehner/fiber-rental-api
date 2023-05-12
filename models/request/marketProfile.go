package requestmodels

type MarketProfile struct {
	Rent             float32 `json:"Rent"`
	PowerPrice       float32 `json:"PowerPrice"`
	Period           uint8   `json:"Period"`
	MinimumPeriod    uint8   `json:"MinimumPeriod"`
	MinimumIncome    float32 `json:"MinimumIncome"`
	Id               string  `json:"Id"`
	HeatPrice        float32 `json:"HeatPrice"`
	Description      string  `json:"Description"`
	AvailabilityDate string  `json:"AvailabilityDate"`
	Deposit          float32 `json:"Deposit"`
}
