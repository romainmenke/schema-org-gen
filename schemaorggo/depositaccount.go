package schemaorggo

import "encoding/json"

// DepositAccount see : https://schema.org/DepositAccount
type DepositAccount struct {
	InvestmentOrDeposit

	typeContext

	// Amount see : https://schema.org/amount
	// The amount of money.
	// types : MonetaryAmount Number
	Amount interface{} `json:"amount,omitempty"`
}

func (v DepositAccount) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "DepositAccount"

	return json.Marshal(v)
}

func (v *DepositAccount) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "DepositAccount"

	return json.Marshal(*v)
}
