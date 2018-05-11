package schemaorg

import "encoding/json"

// InvestmentOrDeposit see : https://schema.org/InvestmentOrDeposit
type InvestmentOrDeposit struct {

typeContext

FinancialProduct

// Amount see : https://schema.org/amount
// The amount of money.
Amount interface{} `json:"amount"` // types : MonetaryAmount Number

}

func (v *InvestmentOrDeposit) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InvestmentOrDeposit"

	return json.Marshal(v)
}
