package requestmodels

type Invite struct {
	Id       string `json:"Id"`
	Property string `json:"Property"`
	Tenant   string `json:"Tenant"`
	Landlord string `json:"Landlord"`
	ValidDue string `json:"ValidDue"`
}
