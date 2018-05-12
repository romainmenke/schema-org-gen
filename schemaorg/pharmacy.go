package schemaorg

import "encoding/json"

// Pharmacy see : https://schema.org/Pharmacy
type Pharmacy struct {

MedicalOrganization

typeContext

// HealthPlanNetworkId see : http://pending.schema.org/healthPlanNetworkId
// Name or unique ID of network. (Networks are often reused across different insurance plans).
HealthPlanNetworkId string `json:"healthPlanNetworkId"`

// IsAcceptingNewPatients see : http://pending.schema.org/isAcceptingNewPatients
// Whether the provider is accepting new patients.
IsAcceptingNewPatients bool `json:"isAcceptingNewPatients"`

// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
// A medical specialty of the provider.
MedicalSpecialty interface{} `json:"medicalSpecialty"`

}

func (v *Pharmacy) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Pharmacy"

	return json.Marshal(v)
}
