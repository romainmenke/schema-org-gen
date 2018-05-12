package schemaorggo

import "encoding/json"

// InsuranceAgency see : https://schema.org/InsuranceAgency
type InsuranceAgency struct {
	FinancialService

	typeContext

	// FeesAndCommissionsSpecification see : https://schema.org/feesAndCommissionsSpecification
	// Description of fees, commissions, and other terms applied either to a class of financial product, or by a financial service organization.
	// types : Text URL
	FeesAndCommissionsSpecification string `json:"feesAndCommissionsSpecification,omitempty"`
}

func (v InsuranceAgency) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "InsuranceAgency"

	return json.Marshal(v)
}

func (v *InsuranceAgency) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "InsuranceAgency"

	return json.Marshal(*v)
}
