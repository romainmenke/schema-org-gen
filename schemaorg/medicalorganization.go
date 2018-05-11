package schemaorg

import "encoding/json"

// MedicalOrganization see : https://schema.org/MedicalOrganization
type MedicalOrganization struct {

typeContext

Organization

// HealthPlanNetworkId see : https://schema.orghttp://pending.schema.org/healthPlanNetworkId
// Name or unique ID of network. (Networks are often reused across different insurance plans).
HealthPlanNetworkId string `json:"healthPlanNetworkId"`

// IsAcceptingNewPatients see : https://schema.orghttp://pending.schema.org/isAcceptingNewPatients
// Whether the provider is accepting new patients.
IsAcceptingNewPatients bool `json:"isAcceptingNewPatients"`

// MedicalSpecialty see : https://schema.orghttp://health-lifesci.schema.org/medicalSpecialty
// A medical specialty of the provider.
MedicalSpecialty interface{} `json:"medicalSpecialty"`

}

func (v *MedicalOrganization) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "MedicalOrganization"

	return json.Marshal(v)
}
