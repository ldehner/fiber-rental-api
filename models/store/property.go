package storemodels

type Property struct {
	Apartment   string `json:"Apartment"`
	Balcony     bool   `json:"Balcony"`
	City        string `json:"City"`
	Country     string `json:"Country"`
	Garage      bool   `json:"Garage"`
	Garden      bool   `json:"Garden"`
	HeatType    uint8  `json:"HeatType"`
	Housenumber string `json:"Housenumber"`
	Id          string `json:"Id"`
	LandlordId  string `json:"Landlord"`
	Rooms       uint8  `json:"Rooms"`
	Size        uint16 `json:"Size"`
	Status      int8   `json:"Status"`
	Street      string `json:"Street"`
	TenantId    string `json:"Tenant"`
	Type        uint8  `json:"Type"`
}
