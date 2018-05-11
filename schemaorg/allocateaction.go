package schemaorg

import "encoding/json"

// AllocateAction see : https://schema.org/AllocateAction
type AllocateAction struct {

typeContext

OrganizeAction

// Purpose see : https://schema.org/purpose
// A goal towards an action is taken. Can be concrete or abstract.
Purpose interface{} `json:"purpose"` // types : MedicalDevicePurpose Thing

}

func (v *AllocateAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AllocateAction"

	return json.Marshal(v)
}
