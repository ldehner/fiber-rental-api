package requestmodels

type User struct {
	Id          string `json:"Id"`
	FirstName   string `json:"First"`
	LastName    string `json:"Last"`
	Phone       string `json:"Phone"`
	Mail        string `json:"Mail"`
	Country     string `json:"Country"`
	City        string `json:"City"`
	Street      string `json:"Street"`
	Housenumber string `json:"Housenumber"`
	Apartment   string `json:"Apartment"`
}
