package schemaorggo

import "encoding/json"

// FinancialService see : https://schema.org/FinancialService
type FinancialService struct {
	LocalBusiness

	typeContext

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	// types : Text URL
	FeesAndCommissionsSpecification string `json:"feesAndCommissionsSpecification,omitempty"`
}

func (v FinancialService) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "FinancialService"

	return json.Marshal(v)
}

func (v *FinancialService) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "FinancialService"

	return json.Marshal(*v)
}
