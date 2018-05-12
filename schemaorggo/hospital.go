package schemaorggo

import "encoding/json"

// Hospital see : https://schema.org/Hospital
type Hospital struct {
	CivicStructure

	typeContext

	// AvailableService see : http://health-lifesci.schema.org/availableService
	// A medical service available from this provider.
	// types : MedicalProcedure MedicalTest MedicalTherapy
	AvailableService interface{} `json:"availableService,omitempty"`

	// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
	// A medical specialty of the provider.
	// types : MedicalSpecialty
	MedicalSpecialty interface{} `json:"medicalSpecialty,omitempty"`
}

func (v Hospital) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Hospital"

	return json.Marshal(v)
}

func (v *Hospital) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Hospital"

	return json.Marshal(*v)
}
