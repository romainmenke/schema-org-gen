package schemaorggo

import "encoding/json"

// AutomatedTeller see : https://schema.org/AutomatedTeller
type AutomatedTeller struct {
	FinancialService

	typeContext

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	// types : Text URL
	FeesAndCommissionsSpecification string `json:"feesAndCommissionsSpecification,omitempty"`
}

func (v AutomatedTeller) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AutomatedTeller"

	return json.Marshal(v)
}

func (v *AutomatedTeller) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "AutomatedTeller"

	return json.Marshal(*v)
}
