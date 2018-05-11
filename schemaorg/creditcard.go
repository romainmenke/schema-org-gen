package schemaorg

import "encoding/json"

// CreditCard see : https://schema.org/CreditCard
type CreditCard struct {

typeContext

LoanOrCredit

// MonthlyMinimumRepaymentAmount see : http://pending.schema.org/monthlyMinimumRepaymentAmount
// The minimum payment is the lowest amount of money that one is required to pay on a credit card statement each month.
MonthlyMinimumRepaymentAmount interface{} `json:"monthlyMinimumRepaymentAmount"` // types : MonetaryAmount Number

}

func (v *CreditCard) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "CreditCard"

	return json.Marshal(v)
}
