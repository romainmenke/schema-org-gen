package schemaorggo

import "encoding/json"

// InvestmentOrDeposit see : https://schema.org/InvestmentOrDeposit
type InvestmentOrDeposit struct {
	FinancialProduct

	typeContext

	// Amount see : https://schema.org/amount
	// The amount of money.
	Amount interface{} `json:"amount,omitempty"` // types : MonetaryAmount Number

}

func (v InvestmentOrDeposit) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InvestmentOrDeposit"

	return json.Marshal(v)
}

func (v *InvestmentOrDeposit) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "InvestmentOrDeposit"

	return json.Marshal(*v)
}
