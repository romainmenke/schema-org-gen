package schemaorg

import "encoding/json"

// DepositAccount see : https://schema.org/DepositAccount
type DepositAccount struct {

typeContext

InvestmentOrDeposit

// Amount see : /amount
// The amount of money.
Amount interface{} `json:"amount"` // types : MonetaryAmount Number

}

func (v *DepositAccount) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DepositAccount"

	return json.Marshal(v)
}
