package schemaorg

import "encoding/json"

// AssignAction see : https://schema.org/AssignAction
type AssignAction struct {

AllocateAction

typeContext

// Purpose see : http://health-lifesci.schema.org/purpose
// A goal towards an action is taken. Can be concrete or abstract.
Purpose interface{} `json:"purpose"` // types : MedicalDevicePurpose Thing

}

func (v *AssignAction) MarshalJSON() ([]byte, error) {
	v.C = "http://schema.org"
	v.T = "AssignAction"

	return json.Marshal(v)
}
