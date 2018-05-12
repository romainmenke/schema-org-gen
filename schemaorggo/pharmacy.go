package schemaorggo

import "encoding/json"

// Pharmacy see : https://schema.org/Pharmacy
type Pharmacy struct {
	MedicalOrganization

	typeContext

	// HealthPlanNetworkId see : http://pending.schema.org/healthPlanNetworkId
	// Name or unique ID of network. (Networks are often reused across different insurance plans).
	HealthPlanNetworkId string `json:"healthPlanNetworkId,omitempty"`

	// IsAcceptingNewPatients see : http://pending.schema.org/isAcceptingNewPatients
	// Whether the provider is accepting new patients.
	IsAcceptingNewPatients bool `json:"isAcceptingNewPatients,omitempty"`

	// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
	// A medical specialty of the provider.
	MedicalSpecialty interface{} `json:"medicalSpecialty,omitempty"`
}

func (v Pharmacy) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Pharmacy"

	return json.Marshal(v)
}

func (v *Pharmacy) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Pharmacy"

	return json.Marshal(*v)
}
