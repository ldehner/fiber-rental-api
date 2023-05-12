package models

type Incident struct {
	Description string `json:"Description"`
	Id          string `json:"Id"`
	Type        uint8  `json:"Type"`
	TenantId    string `json:"Tenant"`
	Status      int8   `json:"Status"`
	PropertyId  string `json:"Property"`
	LandlordId  string `json:"Landlord"`
}
