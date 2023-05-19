package requestmodels

type Incident struct {
	Description string `json:"Description"`
	Type        uint8  `json:"Type"`
	TenantId    string `json:"Tenant"`
	Status      int8   `json:"Status"`
	LandlordId  string `json:"Landlord"`
}
