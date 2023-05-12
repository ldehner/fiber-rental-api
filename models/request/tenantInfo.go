package requestmodels

type TenantInfo struct {
	CriminalRecord bool    `json:"CriminalRecord"`
	Income         float32 `json:"Income"`
	IncomeProof    bool    `json:"IncomeProof"`
}
