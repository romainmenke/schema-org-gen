package schemaorg

import "encoding/json"

// Hospital see : https://schema.org/Hospital
type Hospital struct {

typeContext

CivicStructure

// AvailableService see : http://health-lifesci.schema.org/availableService
// A medical service available from this provider.
AvailableService interface{} `json:"availableService"` // types : MedicalProcedure MedicalTest MedicalTherapy

// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
// A medical specialty of the provider.
MedicalSpecialty interface{} `json:"medicalSpecialty"`

}

func (v *Hospital) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Hospital"

	return json.Marshal(v)
}
