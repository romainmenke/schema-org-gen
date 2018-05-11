package schemaorg

import "encoding/json"

// InsuranceAgency see : https://schema.org/InsuranceAgency
type InsuranceAgency struct {

typeContext

FinancialService

// FeesAndCommissionsSpecification see : /feesAndCommissionsSpecification
// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification"` // types : Text URL

}

func (v *InsuranceAgency) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InsuranceAgency"

	return json.Marshal(v)
}
