package schemaorg

import "encoding/json"

// FinancialService see : https://schema.org/FinancialService
type FinancialService struct {

typeContext

LocalBusiness

// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification"` // types : Text URL

}

func (v *FinancialService) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FinancialService"

	return json.Marshal(v)
}
