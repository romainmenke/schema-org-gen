package schemaorg

import "encoding/json"

// AccountingService see : https://schema.org/AccountingService
type AccountingService struct {

FinancialService

typeContext

// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
FeesAndCommissionsSpecification interface{} `json:"feesAndCommissionsSpecification"` // types : Text URL

}

func (v *AccountingService) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AccountingService"

	return json.Marshal(v)
}
