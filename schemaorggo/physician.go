package schemaorggo

import "encoding/json"

// Physician see : https://schema.org/Physician
type Physician struct {
	MedicalOrganization

	typeContext

	// AvailableService see : http://health-lifesci.schema.org/availableService
	// A medical service available from this provider.
	// types : MedicalProcedure MedicalTest MedicalTherapy
	AvailableService interface{} `json:"availableService,omitempty"`

	// HospitalAffiliation see : http://health-lifesci.schema.org/hospitalAffiliation
	// A hospital with which the physician or office is affiliated.
	// types : Hospital
	HospitalAffiliation *Hospital `json:"hospitalAffiliation,omitempty"`

	// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
	// A medical specialty of the provider.
	// types : MedicalSpecialty
	MedicalSpecialty interface{} `json:"medicalSpecialty,omitempty"`
}

func (v Physician) MarshalJSONWithTypeContext() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Physician"

	return json.Marshal(v)
}

func (v *Physician) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	v.C = "http://schema.org"
	v.T = "Physician"

	return json.Marshal(*v)
}
