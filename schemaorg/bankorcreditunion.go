package schemaorg

import "encoding/json"

// BankOrCreditUnion see : https://schema.org/BankOrCreditUnion
type BankOrCreditUnion struct {

FinancialService

typeContext

// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification"` // types : Text URL

}

func (v *BankOrCreditUnion) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "BankOrCreditUnion"

	return json.Marshal(v)
}
