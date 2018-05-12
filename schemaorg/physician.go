package schemaorg

import "encoding/json"

// Physician see : https://schema.org/Physician
type Physician struct {

MedicalOrganization

typeContext

// AvailableService see : http://health-lifesci.schema.org/availableService
// A medical service available from this provider.
AvailableService interface{} `json:"availableService"` // types : MedicalProcedure MedicalTest MedicalTherapy

// HospitalAffiliation see : http://health-lifesci.schema.org/hospitalAffiliation
// A hospital with which the physician or office is affiliated.
HospitalAffiliation *Hospital `json:"hospitalAffiliation"`

// MedicalSpecialty see : http://health-lifesci.schema.org/medicalSpecialty
// A medical specialty of the provider.
MedicalSpecialty interface{} `json:"medicalSpecialty"`

}

func (v *Physician) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "Physician"

	return json.Marshal(v)
}
