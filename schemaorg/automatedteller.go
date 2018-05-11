package schemaorg

import "encoding/json"

// AutomatedTeller see : https://schema.org/AutomatedTeller
type AutomatedTeller struct {

typeContext

FinancialService

// FeesAndCommissionsSpecification see : /feesAndCommissionsSpecification
// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification"` // types : Text URL

}

func (v *AutomatedTeller) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AutomatedTeller"

	return json.Marshal(v)
}
