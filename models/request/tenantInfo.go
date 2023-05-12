package requestmodels

type TenantInfo struct {
	Id             string  `json:"Id"`
	CriminalRecord bool    `json:"CriminalRecord"`
	Income         float32 `json:"Income"`
	IncomeProof    bool    `json:"IncomeProof"`
}
