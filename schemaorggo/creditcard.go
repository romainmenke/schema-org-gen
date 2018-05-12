package schemaorggo

import "encoding/json"

// CreditCard see : https://schema.org/CreditCard
type CreditCard struct {
	LoanOrCredit

	typeContext

	// MonthlyMinimumRepaymentAmount see : http://pending.schema.org/monthlyMinimumRepaymentAmount
	// The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
	MonthlyMinimumRepaymentAmount interface{} `json:"monthlyMinimumRepaymentAmount,omitempty"` // types : MonetaryAmount Number

}

func (v CreditCard) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CreditCard"

	return json.Marshal(v)
}

func (v *CreditCard) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "CreditCard"

	return json.Marshal(*v)
}
