package schemaorggo

import "encoding/json"

// AccountingService see : https://schema.org/AccountingService
type AccountingService struct {
	FinancialService

	typeContext

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	FeesAndCommissionsSpecification string `json:"feesAndCommissionsSpecification,omitempty"` // types : Text URL

}

func (v AccountingService) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AccountingService"

	return json.Marshal(v)
}

func (v *AccountingService) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AccountingService"

	return json.Marshal(*v)
}
