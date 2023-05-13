package requestmodels

type Invite struct {
	Property string `json:"Property"`
	Tenant   string `json:"Tenant"`
	Landlord string `json:"Landlord"`
}
