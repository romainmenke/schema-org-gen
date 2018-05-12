package schemaorggo

import "encoding/json"

// MedicalOrganization see : https://schema.org/MedicalOrganization
type MedicalOrganization struct {
	Organization

	typeContext

	// HealthPlanNetworkId see : http://pending.schema.org/healthPlanNetworkId
	// Name or unique ID of network. (Networks are often reused across different insurance plans).
	HealthPlanNetworkId string `json:"healthPlanNetworkId,omitempty"` // types : Text

	// IsAcceptingNewPatients see : http://pending.schema.org/isAcceptingNewPatients
	// Whether the provider is accepting new patients.
	IsAcceptingNewPatients bool `json:"isAcceptingNewPatients,omitempty"` // types : Boolean

	// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
	// A medical specialty of the provider.
	MedicalSpecialty interface{} `json:"medicalSpecialty,omitempty"` // types : MedicalSpecialty

}

func (v MedicalOrganization) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MedicalOrganization"

	return json.Marshal(v)
}

func (v *MedicalOrganization) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "MedicalOrganization"

	return json.Marshal(*v)
}
