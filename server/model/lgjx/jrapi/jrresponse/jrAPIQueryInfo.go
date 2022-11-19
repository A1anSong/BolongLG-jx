package jrresponse

type JRAPIQueryInfo struct {
	OrderNo             *string  `json:"orderNo"`
	ElogNo              *string  `json:"elogNo"`
	ProductNo           *string  `json:"productNo"`
	ProductType         *int64   `json:"productType"`
	ProductRate         *float64 `json:"productRate"`
	InsuranceName       *string  `json:"insuranceName"`
	InsuranceCreditCode *string  `json:"insuranceCreditCode"`
	EloOutDate          *string  `json:"eloOutDate"`
	EloUrl              *string  `json:"eloUrl"`
	EloEncryptUrl       *string  `json:"eloEncryptUrl"`
	TenderDeposit       *float64 `json:"tenderDeposit"`
	InsureStartDate     *string  `json:"insureStartDate"`
	InsureEndDate       *string  `json:"insureEndDate"`
	InsureDay           *int64   `json:"insureDay"`
	ValidateCode        *string  `json:"validateCode"`
}
